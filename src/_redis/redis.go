package _redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/junyang7/go-common/src/_as"
	"github.com/junyang7/go-common/src/_codeMessage"
	"github.com/junyang7/go-common/src/_interceptor"
	"math/rand"
	"time"
)

type Connection struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Database string `json:"database"`
	Username string `json:"username"`
	Password string `json:"password"`
}
type Group struct {
	Count      int           `json:"count"`
	Connection []*Connection `json:"connection"`
}
type Cluster struct {
	Master *Group `json:"master"`
	Slaver *Group `json:"slaver"`
}
type Conf struct {
	Count   int        `json:"count"`
	Cluster []*Cluster `json:"cluster"`
}

var redisConf = map[string]*Conf{}
var redisPool = map[string]*redis.Client{}

func Initialize(conf map[string]*Conf) {
	redisConf = conf
	for _, cluster := range conf {
		for _, group := range cluster.Cluster {
			for _, connection := range group.Master.Connection {
				openAndSaveToPool(connection)
			}
			for _, connection := range group.Slaver.Connection {
				openAndSaveToPool(connection)
			}
		}
	}
}

func openAndSaveToPool(connection *Connection) {
	dsn := getDsn(connection)
	if _, ok := redisPool[dsn]; ok {
		return
	}
	p := redis.NewClient(&redis.Options{
		Addr:     connection.Host + ":" + connection.Port,
		DB:       _as.Int(connection.Database),
		Username: connection.Username,
		Password: connection.Password,
	})
	redisPool[dsn] = p
}
func getDsn(connection *Connection) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", connection.Username, connection.Password, connection.Host, connection.Port, connection.Database)
}

type Redis struct {
	baseDatabase  string
	shard         bool
	databaseIndex int
	master        bool
	dsn           string
}

func New(baseDatabase string) *Redis {
	return &Redis{
		baseDatabase: baseDatabase,
	}
}
func (this *Redis) DatabaseIndex(databaseIndex int) *Redis {
	this.databaseIndex = databaseIndex
	return this
}
func (this *Redis) Shard() *Redis {
	this.shard = true
	return this
}
func (this *Redis) Master() *Redis {
	this.master = true
	return this
}
func (this *Redis) getPool() *redis.Client {
	if "" == this.dsn {
		if !this.shard {
			this.databaseIndex = 0
		}
		database, ok := redisConf[this.baseDatabase]
		_interceptor.Insure(ok).
			CodeMessage(_codeMessage.ErrMapKeyNotExists).
			Data(map[string]interface{}{"database": this.baseDatabase}).
			Do()
		_interceptor.Insure(this.databaseIndex < database.Count).
			CodeMessage(_codeMessage.ErrSliceOutOfIndex).
			Data(map[string]interface{}{"database": this.baseDatabase, "index": this.databaseIndex}).
			Do()
		cluster := database.Cluster[this.databaseIndex]
		var group *Group
		if this.master {
			group = cluster.Master
		} else {
			group = cluster.Slaver
		}
		r := 0
		if count := group.Count; count > 1 {
			rand.Seed(time.Now().Unix())
			r = rand.Intn(group.Count - 1)
		}
		this.dsn = getDsn(group.Connection[r])
	}
	return redisPool[this.dsn]
}
func (this *Redis) getCtx() context.Context {
	return context.Background()
}

// generic

// Exists
// 1.0.0
func (this *Redis) Exists(key string) bool {
	res := this.getPool().Exists(this.getCtx(), key).Val()
	return res > 0
}

// Persist
// 2.2.0
func (this *Redis) Persist(key string) bool {
	res := this.getPool().Persist(this.getCtx(), key).Val()
	return res
}

// Expire
// 1.0.0
func (this *Redis) Expire(key string, seconds int) bool {
	res := this.getPool().Expire(this.getCtx(), key, time.Second*time.Duration(seconds)).Val()
	return res
}

// PExpire
// 2.6.0
func (this *Redis) PExpire(key string, milliseconds int) bool {
	res := this.getPool().Expire(this.getCtx(), key, time.Millisecond*time.Duration(milliseconds)).Val()
	return res
}

// Ttl
// 1.0.0
func (this *Redis) Ttl(key string) int {
	res := this.getPool().TTL(this.getCtx(), key).Val()
	if res > 0 {
		return int(res / time.Second)
	}
	return 0
}

// PTtl
// 2.6.0
func (this *Redis) PTtl(key string) int {
	res := this.getPool().TTL(this.getCtx(), key).Val()
	if res > 0 {
		return int(res / time.Millisecond)
	}
	return 0
}

// Scan
// 2.8.0
func (this *Redis) Scan(cursor int, match string, count int, redisType string) ([]string, int) {
	resKeys, resCursor := this.getPool().ScanType(this.getCtx(), uint64(cursor), match, int64(count), redisType).Val()
	return resKeys, int(resCursor)
}

// Type
// 1.0.0
func (this *Redis) Type(key string) string {
	res := this.getPool().Type(this.getCtx(), key).Val()
	return res
}

// Unlink
// 4.0.0
func (this *Redis) Unlink(key ...string) int {
	res := this.getPool().Unlink(this.getCtx(), key...).Val()
	return int(res)
}

// server

// DbSize
// 1.0.0
func (this *Redis) DbSize() int {
	res := this.getPool().DBSize(this.getCtx()).Val()
	return int(res)
}

// Info
// 1.0.0
func (this *Redis) Info() string {
	res := this.getPool().Info(this.getCtx()).Val()
	return res
}

// string

// Decr
// 1.0.0
func (this *Redis) Decr(key string) int {
	res := this.getPool().Decr(this.getCtx(), key).Val()
	return int(res)
}

// DecrBy
// 1.0.0
func (this *Redis) DecrBy(key string, decrement int) int {
	res := this.getPool().DecrBy(this.getCtx(), key, int64(decrement)).Val()
	return int(res)
}

// Get
// 1.0.0
func (this *Redis) Get(key string) string {
	res := this.getPool().Get(this.getCtx(), key).Val()
	return res
}

// Incr
// 1.0.0
func (this *Redis) Incr(key string) int {
	res := this.getPool().Incr(this.getCtx(), key).Val()
	return int(res)
}

// IncrBy
// 1.0.0
func (this *Redis) IncrBy(key string, increment int) int {
	res := this.getPool().IncrBy(this.getCtx(), key, int64(increment)).Val()
	return int(res)
}

// Set
// 1.0.0
// Starting with Redis version 2.6.12: Added the EX, PX, NX and XX options.
func (this *Redis) Set(key string, value interface{}) bool {
	res := this.getPool().Set(this.getCtx(), key, value, 0).Val()
	return "OK" == res
}

// SetNx
// 1.0.0
func (this *Redis) SetNx(key string, value interface{}) bool {
	res := this.getPool().SetNX(this.getCtx(), key, value, 0).Val()
	return res
}

// SetEx
// 2.0.0
func (this *Redis) SetEx(key string, seconds int, value interface{}) bool {
	res := this.getPool().SetEX(this.getCtx(), key, value, time.Second*time.Duration(seconds)).Val()
	return "OK" == res
}

// PSetEx
// 2.6.0
func (this *Redis) PSetEx(key string, milliseconds int, value interface{}) bool {
	res := this.getPool().SetEX(this.getCtx(), key, value, time.Millisecond*time.Duration(milliseconds)).Val()
	return "OK" == res
}

// list

// LLen
// 1.0.0
func (this *Redis) LLen(key string) int {
	res := this.getPool().LLen(this.getCtx(), key).Val()
	return _as.Int(res)
}

// LIndex
// 1.0.0
func (this *Redis) LIndex(key string, index int) string {
	res := this.getPool().LIndex(this.getCtx(), key, int64(index)).Val()
	return res
}

// LRange
// 1.0.0
func (this *Redis) LRange(key string, start int, stop int) []string {
	res := this.getPool().LRange(this.getCtx(), key, int64(start), int64(stop)).Val()
	return res
}

// LRem
// 1.0.0
func (this *Redis) LRem(key string, count int, element interface{}) int {
	res := this.getPool().LRem(this.getCtx(), key, int64(count), element).Val()
	return int(res)
}

// RPush
// 1.0.0
func (this *Redis) RPush(key string, element ...interface{}) int {
	res := this.getPool().RPush(this.getCtx(), key, element...).Val()
	return int(res)
}

// LPush
// 1.0.0
func (this *Redis) LPush(key string, element ...interface{}) int {
	res := this.getPool().LPush(this.getCtx(), key, element...).Val()
	return int(res)
}

// RPushX
// 2.2.0
func (this *Redis) RPushX(key string, element ...interface{}) int {
	res := this.getPool().RPushX(this.getCtx(), key, element...).Val()
	return int(res)
}

// LPushX
// 2.2.0
func (this *Redis) LPushX(key string, element ...interface{}) int {
	res := this.getPool().LPushX(this.getCtx(), key, element...).Val()
	return int(res)
}

// RPop
// 1.0.0
func (this *Redis) RPop(key string, count int) []string {
	res := this.getPool().RPopCount(this.getCtx(), key, count).Val()
	return res
}

// LPop
// 1.0.0
func (this *Redis) LPop(key string, count int) []string {
	res := this.getPool().LPopCount(this.getCtx(), key, count).Val()
	return res
}

// BRPop
// 2.0.0
func (this *Redis) BRPop(seconds int, key ...string) []string {
	res := this.getPool().BRPop(this.getCtx(), time.Second*time.Duration(seconds), key...).Val()
	return res
}

// BLPop
// 2.0.0
func (this *Redis) BLPop(seconds int, key ...string) []string {
	res := this.getPool().BLPop(this.getCtx(), time.Second*time.Duration(seconds), key...).Val()
	return res
}

// hash

// HSet
// 2.0.0
func (this *Redis) HSet(key string, parameter ...interface{}) int {
	res := this.getPool().HSet(this.getCtx(), key, parameter...).Val()
	return int(res)
}

// HGetAll
// HGetAll
func (this *Redis) HGetAll(key string) map[string]string {
	res := this.getPool().HGetAll(this.getCtx(), key).Val()
	return res
}

// HGet
// 2.0.0
func (this *Redis) HGet(key string, field string) string {
	res := this.getPool().HGet(this.getCtx(), key, field).Val()
	return res
}

// HKeys
// 2.0.0
func (this *Redis) HKeys(key string) []string {
	res := this.getPool().HKeys(this.getCtx(), key).Val()
	return res
}

// HVals
// 2.0.0
func (this *Redis) HVals(key string) []string {
	res := this.getPool().HVals(this.getCtx(), key).Val()
	return res
}

// HDel
// 2.0.0
func (this *Redis) HDel(key string, field []string) int {
	res := this.getPool().HDel(this.getCtx(), key, field...).Val()
	return int(res)
}

// HExists
// 2.0.0
func (this *Redis) HExists(key string, field string) bool {
	res := this.getPool().HExists(this.getCtx(), key, field).Val()
	return res
}

// HLen
// 2.0.0
func (this *Redis) HLen(key string) int {
	res := this.getPool().HLen(this.getCtx(), key).Val()
	return int(res)
}

// HMGet
// 2.0.0
func (this *Redis) HMGet(key string, field ...string) []string {
	res := this.getPool().HMGet(this.getCtx(), key, field...).Val()
	resString := make([]string, len(res))
	for k, v := range res {
		resString[k] = _as.String(v)
	}
	return resString
}

// HIncrBy
// 2.0.0
func (this *Redis) HIncrBy(key string, field string, increment int) int {
	res := this.getPool().HIncrBy(this.getCtx(), key, field, int64(increment)).Val()
	return int(res)
}

// HScan
// 2.8.0
func (this *Redis) HScan(key string, cursor int, match string, count int) ([]string, int) {
	resKeys, resCursor := this.getPool().HScan(this.getCtx(), key, uint64(cursor), match, int64(count)).Val()
	return resKeys, int(resCursor)
}

// HSetNx
// 2.0.0
func (this *Redis) HSetNx(key string, field string, value interface{}) bool {
	res := this.getPool().HSetNX(this.getCtx(), key, field, value).Val()
	return res
}

// set

// SAdd
// 1.0.0
func (this *Redis) SAdd(key string, member ...interface{}) int {
	res := this.getPool().SAdd(this.getCtx(), key, member...).Val()
	return int(res)
}

// SCard
// 1.0.0
func (this *Redis) SCard(key string) int {
	res := this.getPool().SCard(this.getCtx(), key).Val()
	return int(res)
}

// SIsMember
// 1.0.0
func (this *Redis) SIsMember(key string, member interface{}) bool {
	res := this.getPool().SIsMember(this.getCtx(), key, member).Val()
	return res
}

// SMIsMember
// 6.2.0
func (this *Redis) SMIsMember(key string, member ...interface{}) []bool {
	res := this.getPool().SMIsMember(this.getCtx(), key, member).Val()
	return res
}

// SMembers
// 1.0.0
func (this *Redis) SMembers(key string) []string {
	res := this.getPool().SMembers(this.getCtx(), key).Val()
	return res
}

// SRem
// 1.0.0
func (this *Redis) SRem(key string, member ...interface{}) int {
	res := this.getPool().SRem(this.getCtx(), key, member).Val()
	return int(res)
}

// SRandMember
// 1.0.0
func (this *Redis) SRandMember(key string, count int) []string {
	res := this.getPool().SRandMemberN(this.getCtx(), key, int64(count)).Val()
	return res
}

// SPop
// 1.0.0
func (this *Redis) SPop(key string, count int) []string {
	res := this.getPool().SPopN(this.getCtx(), key, int64(count)).Val()
	return res
}

// SScan
// 2.8.0
func (this *Redis) SScan(key string, cursor int, match string, count int) ([]string, int) {
	resKeys, resCursor := this.getPool().SScan(this.getCtx(), key, uint64(cursor), match, int64(count)).Val()
	return resKeys, int(resCursor)
}

// SDiff
// 1.0.0
func (this *Redis) SDiff(key ...string) []string {
	res := this.getPool().SDiff(this.getCtx(), key...).Val()
	return res
}

// SDiffStore
// 1.0.0
func (this *Redis) SDiffStore(destination string, key ...string) int {
	res := this.getPool().SDiffStore(this.getCtx(), destination, key...).Val()
	return int(res)
}

// SInter
// 1.0.0
func (this *Redis) SInter(key ...string) []string {
	res := this.getPool().SDiff(this.getCtx(), key...).Val()
	return res
}

// SInterStore
// 1.0.0
func (this *Redis) SInterStore(destination string, key ...string) int {
	res := this.getPool().SInterStore(this.getCtx(), destination, key...).Val()
	return int(res)
}

// SUnion
// 1.0.0
func (this *Redis) SUnion(key []string) []string {
	res := this.getPool().SUnion(this.getCtx(), key...).Val()
	return res
}

// SUnionStore
// 1.0.0
func (this *Redis) SUnionStore(destination string, key ...string) int {
	res := this.getPool().SUnionStore(this.getCtx(), destination, key...).Val()
	return int(res)
}

// SMove
// 1.0.0
func (this *Redis) SMove(source string, destination string, member interface{}) bool {
	res := this.getPool().SMove(this.getCtx(), source, destination, member).Val()
	return res
}

// zset

// ZAdd
// 1.2.0
func (this *Redis) ZAdd(key string, parameter ...interface{}) int {
	member := []*redis.Z{}
	for i, j := 0, len(parameter); i < j; i += 2 {
		member = append(member, &redis.Z{
			Score:  _as.Float64(parameter[i]),
			Member: parameter[i+1],
		})
	}
	res := this.getPool().ZAdd(this.getCtx(), key, member...).Val()
	return int(res)
}

// ZCard
// 1.2.0
func (this *Redis) ZCard(key string) int {
	res := this.getPool().ZCard(this.getCtx(), key).Val()
	return _as.Int(res)
}

// ZRange
// 1.2.0
func (this *Redis) ZRange(key string, start int, stop int) []string {
	res := this.getPool().ZRange(this.getCtx(), key, int64(start), int64(stop)).Val()
	return res
}

// ZRangeByScore
// 1.0.5-6.2.0
func (this *Redis) ZRangeByScore(key string, min int, max int) []string {
	res := this.getPool().ZRangeByScore(this.getCtx(), key, &redis.ZRangeBy{Min: _as.String(min), Max: _as.String(max)}).Val()
	return res
}

// ZRem
// 1.2.0
func (this *Redis) ZRem(key string, member ...interface{}) int {
	res := this.getPool().ZRem(this.getCtx(), key, member).Val()
	return int(res)
}

// ZScan
// 2.8.0
func (this *Redis) ZScan(key string, cursor int, match string, count int) ([]string, int) {
	resKeys, resCursor := this.getPool().ZScan(this.getCtx(), key, uint64(cursor), match, int64(count)).Val()
	return resKeys, int(resCursor)
}

// bitmap

// SetBit
// 2.2.0
func (this *Redis) SetBit(key string, offset int, value int) int {
	res := this.getPool().SetBit(this.getCtx(), key, int64(offset), value).Val()
	return int(res)
}

// GetBit
// 2.2.0
func (this *Redis) GetBit(key string, offset int) int {
	res := this.getPool().GetBit(this.getCtx(), key, int64(offset)).Val()
	return int(res)
}

// BitCount
// 2.6.0
func (this *Redis) BitCount(key string, start int, end int) int {
	res := this.getPool().BitCount(this.getCtx(), key, &redis.BitCount{Start: int64(start), End: int64(end)}).Val()
	return int(res)
}

// Eval
// 2.6.0
func (this *Redis) Eval(script string, key []string, arg ...interface{}) interface{} {
	res := this.getPool().Eval(this.getCtx(), script, key, arg...).Val()
	return res
}

// EvalSha
// 2.6.0
func (this *Redis) EvalSha(sha1 string, key []string, arg ...interface{}) interface{} {
	res := this.getPool().EvalSha(this.getCtx(), sha1, key, arg...).Val()
	return res
}

// ScriptExists
// 2.6.0
func (this *Redis) ScriptExists(sha1 ...string) []bool {
	res := this.getPool().ScriptExists(this.getCtx(), sha1...).Val()
	return res
}

// ScriptLoad
// 2.6.0
func (this *Redis) ScriptLoad(script string) string {
	res := this.getPool().ScriptLoad(this.getCtx(), script).Val()
	return res
}

// ScriptFlush
// 2.6.0
func (this *Redis) ScriptFlush() bool {
	res := this.getPool().ScriptFlush(this.getCtx()).Val()
	return "OK" == res
}
