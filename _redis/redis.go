package _redis

import (
	"context"
	"fmt"
	"github.com/junyang7/go-common/_as"
	"github.com/junyang7/go-common/_conf"
	"github.com/junyang7/go-common/_interceptor"
	"github.com/junyang7/go-common/_is"
	"github.com/junyang7/go-common/_json"
	"github.com/redis/go-redis/v9"
	"math/rand"
	"net/url"
	"sync"
	"time"
)

type Machine struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Database string `json:"database"`
	Username string `json:"username"`
	Password string `json:"password"`
}
type Business struct {
	Master []*Machine `json:"master"`
	Slaver []*Machine `json:"slaver"`
}

var poolDict = map[string][]*redis.Client{}
var m = sync.RWMutex{}

func getPoolDictName(business string, master bool) string {
	poolDictName := business + "."
	if master {
		poolDictName += "master"
	} else {
		poolDictName += "slaver"
	}
	return poolDictName
}
func getDsn(machine *Machine) string {
	var dsn string
	dsn = fmt.Sprintf("redis://%s@%s:%s/%s", url.UserPassword(machine.Username, machine.Password), machine.Host, machine.Port, machine.Database)
	return dsn
}
func open(machine *Machine) (pool *redis.Client) {
	dsn := getDsn(machine)
	opt, err := redis.ParseURL(dsn)
	if nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	opt.PoolSize = 50
	opt.MinIdleConns = 5
	pool = redis.NewClient(opt)
	return pool
}
func Load() {
	raw := _conf.Get("redis").Value()
	var businessList map[string]*Business
	_json.Decode(_json.Encode(raw), &businessList)
	if len(businessList) == 0 {
		return
	}
	for business, ms := range businessList {
		for _, machine := range ms.Master {
			poolDictName := getPoolDictName(business, true)
			poolDict[poolDictName] = append(poolDict[poolDictName], open(machine))
		}
		for _, machine := range ms.Slaver {
			poolDictName := getPoolDictName(business, false)
			poolDict[poolDictName] = append(poolDict[poolDictName], open(machine))
		}
	}
}

type Redis struct {
	business          string
	masterMachineList []*Machine
	slaverMachineList []*Machine
	master            bool
	context           context.Context
}

func New() *Redis {
	return &Redis{
		context:  context.Background(),
		business: "default",
	}
}
func (this *Redis) Context(context context.Context) *Redis {
	this.context = context
	return this
}
func (this *Redis) Business(business string) *Redis {
	this.business = business
	return this
}
func (this *Redis) Machine(machine *Machine) *Redis {
	this.masterMachineList = append(this.masterMachineList, machine)
	this.slaverMachineList = append(this.slaverMachineList, machine)
	return this
}
func (this *Redis) MasterMachine(machine *Machine) *Redis {
	this.masterMachineList = append(this.masterMachineList, machine)
	return this
}
func (this *Redis) SlaverMachine(machine *Machine) *Redis {
	this.slaverMachineList = append(this.slaverMachineList, machine)
	return this
}
func (this *Redis) Master(master bool) *Redis {
	this.master = master
	return this
}
func (this *Redis) getMaster() bool {
	return this.master
}
func (this *Redis) getBusiness() string {
	return this.business
}
func (this *Redis) getMasterMachineList() []*Machine {
	return this.masterMachineList
}
func (this *Redis) getSlaverMachineList() []*Machine {
	return this.slaverMachineList
}
func (this *Redis) getPool() *redis.Client {
	business := this.getBusiness()
	master := this.getMaster()
	poolDictName := getPoolDictName(business, master)
	var pool *redis.Client
	var poolList []*redis.Client
	var ok bool
	m.RLock()
	poolList, ok = poolDict[poolDictName]
	m.RUnlock()
	if ok {
		r := rand.Intn(len(poolList))
		return poolList[r]
	}
	m.Lock()
	defer m.Unlock()
	poolList, ok = poolDict[poolDictName]
	if ok {
		r := rand.Intn(len(poolList))
		return poolList[r]
	}
	var machineList []*Machine
	if master {
		machineList = this.getMasterMachineList()
	} else {
		machineList = this.getSlaverMachineList()
	}
	if len(machineList) > 0 {
		r := rand.Intn(len(machineList))
		machine := machineList[r]
		pool = open(machine)
		poolDict[poolDictName] = append(poolDict[poolDictName], open(machine))
		return pool
	}
	_interceptor.Insure(false).Message("没有找到相关配置").Do()
	return nil
}
func (this *Redis) getContext() context.Context {
	if _is.Empty(this.context) {
		this.context = context.Background()
	}
	return this.context
}
func (this *Redis) Do(parameter ...interface{}) *redis.Cmd {
	res := this.getPool().Do(this.getContext(), parameter...)
	err := res.Err()
	if nil != err && redis.Nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return res
}
func (this *Redis) Unlink(keyList ...string) int64 {
	res, err := this.getPool().Unlink(this.getContext(), keyList...).Result()
	if nil != err && redis.Nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return res
}
func (this *Redis) Del(keyList ...string) int64 {
	res, err := this.getPool().Del(this.getContext(), keyList...).Result()
	if nil != err && redis.Nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return res
}
func (this *Redis) Dump(key string) string {
	res, err := this.getPool().Dump(this.getContext(), key).Result()
	if nil != err && redis.Nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return res
}
func (this *Redis) Exists(key string) int64 {
	res, err := this.getPool().Exists(this.getContext(), key).Result()
	if nil != err && redis.Nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return res
}
func (this *Redis) Expire(key string, expiration time.Duration) bool {
	res, err := this.getPool().Expire(this.getContext(), key, expiration).Result()
	if nil != err && redis.Nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return res
}
func (this *Redis) ExpireAt(key string, t time.Time) bool {
	res, err := this.getPool().ExpireAt(this.getContext(), key, t).Result()
	if nil != err && redis.Nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return res
}
func (this *Redis) Keys(pattern string) []string {
	res, err := this.getPool().Keys(this.getContext(), pattern).Result()
	if nil != err && redis.Nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return res
}
func (this *Redis) Migrate(host string, port string, key string, database int, timeout time.Duration) string {
	res, err := this.getPool().Migrate(this.getContext(), host, port, key, database, timeout).Result()
	if nil != err && redis.Nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return res
}
func (this *Redis) Move(key string, database int) bool {
	res, err := this.getPool().Move(this.getContext(), key, database).Result()
	if nil != err && redis.Nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return res
}
func (this *Redis) Persist(key string) bool {
	res, err := this.getPool().Persist(this.getContext(), key).Result()
	if nil != err && redis.Nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return res
}
func (this *Redis) PExpire(key string, expiration time.Duration) bool {
	res, err := this.getPool().PExpire(this.getContext(), key, expiration).Result()
	if nil != err && redis.Nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return res
}
func (this *Redis) PExpireAt(key string, t time.Time) bool {
	res, err := this.getPool().PExpireAt(this.getContext(), key, t).Result()
	if nil != err && redis.Nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return res
}
func (this *Redis) PTtl(key string) time.Duration {
	res, err := this.getPool().PTTL(this.getContext(), key).Result()
	if nil != err && redis.Nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return res
}
func (this *Redis) RandomKey() string {
	res, err := this.getPool().RandomKey(this.getContext()).Result()
	if nil != err && redis.Nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return res
}
func (this *Redis) Rename(old string, new string) string {
	res, err := this.getPool().Rename(this.getContext(), old, new).Result()
	if nil != err && redis.Nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return res
}
func (this *Redis) RenameNx(old string, new string) bool {
	res, err := this.getPool().RenameNX(this.getContext(), old, new).Result()
	if nil != err && redis.Nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return res
}
func (this *Redis) Restore(key string, ttl time.Duration, value string) string {
	res, err := this.getPool().Restore(this.getContext(), key, ttl, value).Result()
	if nil != err && redis.Nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return res
}
func (this *Redis) RestoreReplace(key string, ttl time.Duration, value string) string {
	res, err := this.getPool().RestoreReplace(this.getContext(), key, ttl, value).Result()
	if nil != err && redis.Nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return res
}
func (this *Redis) Ttl(key string) time.Duration {
	res, err := this.getPool().TTL(this.getContext(), key).Result()
	if nil != err && redis.Nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return res
}
func (this *Redis) Type(key string) string {
	res, err := this.getPool().Type(this.getContext(), key).Result()
	if nil != err && redis.Nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return res
}
func (this *Redis) Scan(cursor uint64, match string, count int64) ([]string, uint64) {
	keyList, cursor, err := this.getPool().Scan(this.getContext(), cursor, match, count).Result()
	if nil != err && redis.Nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return keyList, cursor
}
func (this *Redis) Append(key string, value string) int64 {
	res, err := this.getPool().Append(this.getContext(), key, value).Result()
	if nil != err && redis.Nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return res
}
func (this *Redis) BitCount(key string, bitCount *redis.BitCount) int64 {
	res, err := this.getPool().BitCount(this.getContext(), key, bitCount).Result()
	if nil != err && redis.Nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return res
}
func (this *Redis) Decr(key string) int64 {
	res, err := this.getPool().Decr(this.getContext(), key).Result()
	if nil != err && redis.Nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return res
}
func (this *Redis) DecrBy(key string, decrement int64) int64 {
	res, err := this.getPool().DecrBy(this.getContext(), key, decrement).Result()
	if nil != err && redis.Nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return res
}
func (this *Redis) Get(key string) string {
	res, err := this.getPool().Get(this.getContext(), key).Result()
	if nil != err && redis.Nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return res
}
func (this *Redis) GetBit(key string, offset int64) int64 {
	res, err := this.getPool().GetBit(this.getContext(), key, offset).Result()
	if nil != err && redis.Nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return res
}
func (this *Redis) GetRange(key string, start int64, end int64) string {
	res, err := this.getPool().GetRange(this.getContext(), key, start, end).Result()
	if nil != err && redis.Nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return res
}
func (this *Redis) GetSet(key string, value interface{}) string {
	res, err := this.getPool().GetSet(this.getContext(), key, value).Result()
	if nil != err && redis.Nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return res
}
func (this *Redis) Incr(key string) int64 {
	res, err := this.getPool().Incr(this.getContext(), key).Result()
	if nil != err && redis.Nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return res
}
func (this *Redis) IncrBy(key string, value int64) int64 {
	res, err := this.getPool().IncrBy(this.getContext(), key, value).Result()
	if nil != err && redis.Nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return res
}
func (this *Redis) IncrByFloat(key string, value float64) float64 {
	res, err := this.getPool().IncrByFloat(this.getContext(), key, value).Result()
	if nil != err && redis.Nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return res
}
func (this *Redis) MGet(keyList ...string) []string {
	res, err := this.getPool().MGet(this.getContext(), keyList...).Result()
	if nil != err && redis.Nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	resString := make([]string, len(res))
	for k, v := range res {
		resString[k] = _as.String(v)
	}
	return resString
}
func (this *Redis) MSet(parameterList ...interface{}) string {
	res, err := this.getPool().MSet(this.getContext(), parameterList...).Result()
	if nil != err && redis.Nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return res
}
func (this *Redis) MSetNx(parameterList ...interface{}) bool {
	res, err := this.getPool().MSetNX(this.getContext(), parameterList...).Result()
	if nil != err && redis.Nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return res
}
func (this *Redis) Set(key string, value interface{}, expiration time.Duration) string {
	res, err := this.getPool().Set(this.getContext(), key, value, expiration).Result()
	if nil != err && redis.Nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return res
}
func (this *Redis) SetBit(key string, offset int64, value int) int64 {
	res, err := this.getPool().SetBit(this.getContext(), key, offset, value).Result()
	if nil != err && redis.Nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return res
}
func (this *Redis) SetEx(key string, seconds int64, value interface{}, expiration time.Duration) string {
	res, err := this.getPool().SetEx(this.getContext(), key, value, expiration).Result()
	if nil != err && redis.Nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return res
}
func (this *Redis) SetNx(key string, value interface{}, expiration time.Duration) bool {
	res, err := this.getPool().SetNX(this.getContext(), key, value, expiration).Result()
	if nil != err && redis.Nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return res
}
func (this *Redis) SetRange(key string, offset int64, value string) int64 {
	res, err := this.getPool().SetRange(this.getContext(), key, offset, value).Result()
	if nil != err && redis.Nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return res
}
func (this *Redis) StrLen(key string) int64 {
	res, err := this.getPool().StrLen(this.getContext(), key).Result()
	if nil != err && redis.Nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return res
}
func (this *Redis) HDel(key string, field string) int64 {
	res, err := this.getPool().HDel(this.getContext(), key, field).Result()
	if nil != err && redis.Nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return res
}
func (this *Redis) HExists(key string, field string) bool {
	res, err := this.getPool().HExists(this.getContext(), key, field).Result()
	if nil != err && redis.Nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return res
}
func (this *Redis) HGet(key string, field string) string {
	res, err := this.getPool().HGet(this.getContext(), key, field).Result()
	if nil != err && redis.Nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return res
}
func (this *Redis) HGetAll(key string) map[string]string {
	res, err := this.getPool().HGetAll(this.getContext(), key).Result()
	if nil != err && redis.Nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return res
}
func (this *Redis) HIncrBy(key string, field string, increment int64) int64 {
	res, err := this.getPool().HIncrBy(this.getContext(), key, field, increment).Result()
	if nil != err && redis.Nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return res
}
func (this *Redis) HIncrByFloat(key string, field string, increment float64) float64 {
	res, err := this.getPool().HIncrByFloat(this.getContext(), key, field, increment).Result()
	if nil != err && redis.Nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return res
}
func (this *Redis) HKeys(key string) []string {
	res, err := this.getPool().HKeys(this.getContext(), key).Result()
	if nil != err && redis.Nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return res
}
func (this *Redis) HLen(key string) int64 {
	res, err := this.getPool().HLen(this.getContext(), key).Result()
	if nil != err && redis.Nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return res
}
func (this *Redis) HMGet(key string, field ...string) []string {
	res, err := this.getPool().HMGet(this.getContext(), key, field...).Result()
	if nil != err && redis.Nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	resString := make([]string, len(res))
	for k, v := range res {
		resString[k] = _as.String(v)
	}
	return resString
}
func (this *Redis) HMSet(key string, parameter ...interface{}) bool {
	res, err := this.getPool().HMSet(this.getContext(), key, parameter...).Result()
	if nil != err && redis.Nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return res
}
func (this *Redis) HSet(key string, parameter ...interface{}) int64 {
	res, err := this.getPool().HSet(this.getContext(), key, parameter...).Result()
	if nil != err && redis.Nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return res
}
func (this *Redis) HSetNX(key string, field string, value interface{}) bool {
	res, err := this.getPool().HSetNX(this.getContext(), key, field, value).Result()
	if nil != err && redis.Nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return res
}
func (this *Redis) HVals(key string) []string {
	res, err := this.getPool().HVals(this.getContext(), key).Result()
	if nil != err && redis.Nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return res
}
func (this *Redis) HScan(key string, cursor uint64, match string, count int64) ([]string, uint64) {
	keyList, cursor, err := this.getPool().HScan(this.getContext(), key, cursor, match, count).Result()
	if nil != err && redis.Nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return keyList, cursor
}
func (this *Redis) BLPop(timeout time.Duration, keyList ...string) []string {
	res, err := this.getPool().BLPop(this.getContext(), timeout, keyList...).Result()
	if nil != err && redis.Nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return res
}
func (this *Redis) BRPop(timeout time.Duration, keyList ...string) []string {
	res, err := this.getPool().BRPop(this.getContext(), timeout, keyList...).Result()
	if nil != err && redis.Nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return res
}
func (this *Redis) BRPopLPush(r string, l string, timeout time.Duration) string {
	res, err := this.getPool().BRPopLPush(this.getContext(), r, l, timeout).Result()
	if nil != err && redis.Nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return res
}
func (this *Redis) LIndex(key string, index int64) string {
	res, err := this.getPool().LIndex(this.getContext(), key, index).Result()
	if nil != err && redis.Nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return res
}
func (this *Redis) LLen(key string) int64 {
	res, err := this.getPool().LLen(this.getContext(), key).Result()
	if nil != err && redis.Nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return res
}
func (this *Redis) LPop(key string) string {
	res, err := this.getPool().LPop(this.getContext(), key).Result()
	if nil != err && redis.Nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return res
}
func (this *Redis) LPush(key string, valueList ...interface{}) int64 {
	res, err := this.getPool().LPush(this.getContext(), key, valueList...).Result()
	if nil != err && redis.Nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return res
}
func (this *Redis) LPushX(key string, valueList ...interface{}) int64 {
	res, err := this.getPool().LPushX(this.getContext(), key, valueList...).Result()
	if nil != err && redis.Nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return res
}
func (this *Redis) LRange(key string, start int64, stop int64) []string {
	res, err := this.getPool().LRange(this.getContext(), key, start, stop).Result()
	if nil != err && redis.Nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return res
}
func (this *Redis) LRem(key string, count int64, value interface{}) int64 {
	res, err := this.getPool().LRem(this.getContext(), key, count, value).Result()
	if nil != err && redis.Nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return res
}
func (this *Redis) LSet(key string, index int64, value interface{}) string {
	res, err := this.getPool().LSet(this.getContext(), key, index, value).Result()
	if nil != err && redis.Nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return res
}
func (this *Redis) LTrim(key string, start int64, stop int64) string {
	res, err := this.getPool().LTrim(this.getContext(), key, start, stop).Result()
	if nil != err && redis.Nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return res
}
func (this *Redis) RPop(key string) string {
	res, err := this.getPool().RPop(this.getContext(), key).Result()
	if nil != err && redis.Nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return res
}
func (this *Redis) RPopLPush(r string, l string) string {
	res, err := this.getPool().RPopLPush(this.getContext(), r, l).Result()
	if nil != err && redis.Nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return res
}
func (this *Redis) RPush(key string, valueList ...interface{}) int64 {
	res, err := this.getPool().RPush(this.getContext(), key, valueList...).Result()
	if nil != err && redis.Nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return res
}
func (this *Redis) RPushX(key string, valueList ...interface{}) int64 {
	res, err := this.getPool().RPushX(this.getContext(), key, valueList...).Result()
	if nil != err && redis.Nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return res
}
func (this *Redis) SAdd(key string, valueList ...interface{}) int64 {
	res, err := this.getPool().SAdd(this.getContext(), key, valueList...).Result()
	if nil != err && redis.Nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return res
}
func (this *Redis) SCard(key string) int64 {
	res, err := this.getPool().SCard(this.getContext(), key).Result()
	if nil != err && redis.Nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return res
}
func (this *Redis) SDiff(keyList ...string) []string {
	res, err := this.getPool().SDiff(this.getContext(), keyList...).Result()
	if nil != err && redis.Nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return res
}
func (this *Redis) SDiffStore(store string, keyList ...string) int64 {
	res, err := this.getPool().SDiffStore(this.getContext(), store, keyList...).Result()
	if nil != err && redis.Nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return res
}
func (this *Redis) SInter(keyList ...string) []string {
	res, err := this.getPool().SInter(this.getContext(), keyList...).Result()
	if nil != err && redis.Nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return res
}
func (this *Redis) SInterStore(store string, keyList ...string) int64 {
	res, err := this.getPool().SInterStore(this.getContext(), store, keyList...).Result()
	if nil != err && redis.Nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return res
}
func (this *Redis) SIsMember(key string, value interface{}) bool {
	res, err := this.getPool().SIsMember(this.getContext(), key, value).Result()
	if nil != err && redis.Nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return res
}
func (this *Redis) SMembers(key string) []string {
	res, err := this.getPool().SMembers(this.getContext(), key).Result()
	if nil != err && redis.Nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return res
}
func (this *Redis) SMove(old string, new string, member interface{}) bool {
	res, err := this.getPool().SMove(this.getContext(), old, new, member).Result()
	if nil != err && redis.Nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return res
}
func (this *Redis) SPop(key string) string {
	res, err := this.getPool().SPop(this.getContext(), key).Result()
	if nil != err && redis.Nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return res
}
func (this *Redis) SRandMember(key string) string {
	res, err := this.getPool().SRandMember(this.getContext(), key).Result()
	if nil != err && redis.Nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return res
}
func (this *Redis) SRem(key string, memberList ...interface{}) int64 {
	res, err := this.getPool().SRem(this.getContext(), key, memberList...).Result()
	if nil != err && redis.Nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return res
}
func (this *Redis) SUnion(keyList ...string) []string {
	res, err := this.getPool().SUnion(this.getContext(), keyList...).Result()
	if nil != err && redis.Nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return res
}
func (this *Redis) SUnionStore(store string, keyList ...string) int64 {
	res, err := this.getPool().SUnionStore(this.getContext(), store, keyList...).Result()
	if nil != err && redis.Nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return res
}
func (this *Redis) SScan(key string, cursor uint64, match string, count int64) ([]string, uint64) {
	keyList, cursor, err := this.getPool().SScan(this.getContext(), key, cursor, match, count).Result()
	if nil != err && redis.Nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return keyList, cursor
}
func (this *Redis) ZAdd(key string, valueList ...interface{}) int64 {
	memberList := []redis.Z{}
	for i, j := 0, len(valueList); i < j; i += 2 {
		memberList = append(memberList, redis.Z{
			Score:  _as.Float64(valueList[i]),
			Member: valueList[i+1],
		})
	}
	res := this.getPool().ZAdd(this.getContext(), key, memberList...).Val()
	return res
}
func (this *Redis) ZCard(key string) int64 {
	res, err := this.getPool().ZCard(this.getContext(), key).Result()
	if nil != err && redis.Nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return res
}
func (this *Redis) ZCount(key string, min string, max string) int64 {
	res, err := this.getPool().ZCount(this.getContext(), key, min, max).Result()
	if nil != err && redis.Nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return res
}
func (this *Redis) ZIncrBy(key string, increment float64, member string) float64 {
	res, err := this.getPool().ZIncrBy(this.getContext(), key, increment, member).Result()
	if nil != err && redis.Nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return res
}
func (this *Redis) ZRange(key string, start int64, stop int64) []string {
	res, err := this.getPool().ZRange(this.getContext(), key, start, stop).Result()
	if nil != err && redis.Nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return res
}
func (this *Redis) ZRangeByScore(key string, opt *redis.ZRangeBy) []string {
	res, err := this.getPool().ZRangeByScore(this.getContext(), key, opt).Result()
	if nil != err && redis.Nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return res
}
func (this *Redis) ZRank(key string, member string) int64 {
	res, err := this.getPool().ZRank(this.getContext(), key, member).Result()
	if nil != err && redis.Nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return res
}
func (this *Redis) ZRem(key string, memberList ...interface{}) int64 {
	res, err := this.getPool().ZRem(this.getContext(), key, memberList...).Result()
	if nil != err && redis.Nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return res
}
func (this *Redis) ZRemRangeByRank(key string, start int64, stop int64) int64 {
	res, err := this.getPool().ZRemRangeByRank(this.getContext(), key, start, stop).Result()
	if nil != err && redis.Nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return res
}
func (this *Redis) ZRemRangeByScore(key string, min string, max string) int64 {
	res, err := this.getPool().ZRemRangeByScore(this.getContext(), key, min, max).Result()
	if nil != err && redis.Nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return res
}
func (this *Redis) ZRevRange(key string, start int64, stop int64) []string {
	res, err := this.getPool().ZRevRange(this.getContext(), key, start, stop).Result()
	if nil != err && redis.Nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return res
}
func (this *Redis) ZRevRangeByScore(key string, opt *redis.ZRangeBy) []string {
	res, err := this.getPool().ZRevRangeByScore(this.getContext(), key, opt).Result()
	if nil != err && redis.Nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return res
}
func (this *Redis) ZRevRank(key string, member string) int64 {
	res, err := this.getPool().ZRevRank(this.getContext(), key, member).Result()
	if nil != err && redis.Nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return res
}
func (this *Redis) ZScore(key string, member string) float64 {
	res, err := this.getPool().ZScore(this.getContext(), key, member).Result()
	if nil != err && redis.Nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return res
}
func (this *Redis) ZScan(key string, cursor uint64, match string, count int64) ([]string, uint64) {
	keyList, cursor, err := this.getPool().ZScan(this.getContext(), key, cursor, match, count).Result()
	if nil != err && redis.Nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return keyList, cursor
}
func (this *Redis) Eval(script string, keyList []string, parameterList ...interface{}) interface{} {
	res, err := this.getPool().Eval(this.getContext(), script, keyList, parameterList...).Result()
	if nil != err && redis.Nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return res
}
func (this *Redis) EvalSha(sha1 string, keyList []string, parameterList ...interface{}) interface{} {
	res, err := this.getPool().EvalSha(this.getContext(), sha1, keyList, parameterList...).Result()
	if nil != err && redis.Nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return res
}
func (this *Redis) ScriptExists(sha1List ...string) []bool {
	res, err := this.getPool().ScriptExists(this.getContext(), sha1List...).Result()
	if nil != err && redis.Nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return res
}
func (this *Redis) ScriptFlush() string {
	res, err := this.getPool().ScriptFlush(this.getContext()).Result()
	if nil != err && redis.Nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return res
}
func (this *Redis) ScriptKill() string {
	res, err := this.getPool().ScriptKill(this.getContext()).Result()
	if nil != err && redis.Nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return res
}
func (this *Redis) ScriptLoad(script string) string {
	res, err := this.getPool().ScriptLoad(this.getContext(), script).Result()
	if nil != err && redis.Nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return res
}
func (this *Redis) BgRewriteAOF() string {
	res, err := this.getPool().BgRewriteAOF(this.getContext()).Result()
	if nil != err && redis.Nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return res
}
func (this *Redis) BgSave() string {
	res, err := this.getPool().BgSave(this.getContext()).Result()
	if nil != err && redis.Nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return res
}
func (this *Redis) ClientGetName() string {
	res, err := this.getPool().ClientGetName(this.getContext()).Result()
	if nil != err && redis.Nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return res
}
func (this *Redis) ClientKill(ipPort string) string {
	res, err := this.getPool().ClientKill(this.getContext(), ipPort).Result()
	if nil != err && redis.Nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return res
}
func (this *Redis) ClientList() string {
	res, err := this.getPool().ClientList(this.getContext()).Result()
	if nil != err && redis.Nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return res
}
func (this *Redis) ConfigSet(label string, value string) string {
	res, err := this.getPool().ConfigSet(this.getContext(), label, value).Result()
	if nil != err && redis.Nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return res
}
func (this *Redis) ConfigGet(label string) map[string]string {
	res, err := this.getPool().ConfigGet(this.getContext(), label).Result()
	if nil != err && redis.Nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return res
}
func (this *Redis) ConfigResetStat() string {
	res, err := this.getPool().ConfigResetStat(this.getContext()).Result()
	if nil != err && redis.Nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return res
}
func (this *Redis) ConfigRewrite() string {
	res, err := this.getPool().ConfigRewrite(this.getContext()).Result()
	if nil != err && redis.Nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return res
}
func (this *Redis) DBSize() int64 {
	res, err := this.getPool().DBSize(this.getContext()).Result()
	if nil != err && redis.Nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return res
}
func (this *Redis) FlushAll() string {
	res, err := this.getPool().FlushAll(this.getContext()).Result()
	if nil != err && redis.Nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return res
}
func (this *Redis) FlushDb() string {
	res, err := this.getPool().FlushDB(this.getContext()).Result()
	if nil != err && redis.Nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return res
}
func (this *Redis) Info() string {
	res, err := this.getPool().Info(this.getContext()).Result()
	if nil != err && redis.Nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return res
}
func (this *Redis) LastSave() int64 {
	res, err := this.getPool().LastSave(this.getContext()).Result()
	if nil != err && redis.Nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return res
}
func (this *Redis) Save() string {
	res, err := this.getPool().Save(this.getContext()).Result()
	if nil != err && redis.Nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return res
}
func (this *Redis) SlowLogGet(num int64) []redis.SlowLog {
	res, err := this.getPool().SlowLogGet(this.getContext(), num).Result()
	if nil != err && redis.Nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return res
}
func (this *Redis) Ping() string {
	res, err := this.getPool().Ping(this.getContext()).Result()
	if nil != err && redis.Nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return res
}
func (this *Redis) Echo(message interface{}) string {
	res, err := this.getPool().Echo(this.getContext(), message).Result()
	if nil != err && redis.Nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return res
}
