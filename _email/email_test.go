package _email

import (
	"github.com/junyang7/go-common/_assert"
	"github.com/junyang7/go-common/_file"
	"testing"
)

func TestNew(t *testing.T) {
	{
		m := New("smtp.example.com", 587, "user", "pass", "from@example.com", "to@example.com")
		_assert.Equal(t, "smtp.example.com", m.host)
		_assert.Equal(t, 587, m.port)
		_assert.Equal(t, "user", m.username)
		_assert.Equal(t, "pass", m.password)
		_assert.Equal(t, "from@example.com", m.from)
		_assert.EqualByList(t, []string{"to@example.com"}, m.toList)
		_assert.Equal(t, "", m.subject)
		_assert.Equal(t, "", m.content)
		_assert.EqualByList(t, []string{}, m.ccList)
		_assert.EqualByList(t, []string{}, m.attachList)
	}
	{
		m := New("smtp.example.com", 587, "user", "pass", "from@example.com", "to1@example.com", "to2@example.com")
		_assert.Equal(t, "smtp.example.com", m.host)
		_assert.Equal(t, 587, m.port)
		_assert.Equal(t, "user", m.username)
		_assert.Equal(t, "pass", m.password)
		_assert.Equal(t, "from@example.com", m.from)
		_assert.EqualByList(t, []string{"to1@example.com", "to2@example.com"}, m.toList)
		_assert.Equal(t, "", m.subject)
		_assert.Equal(t, "", m.content)
		_assert.EqualByList(t, []string{}, m.ccList)
		_assert.EqualByList(t, []string{}, m.attachList)
	}
}
func TestSubject(t *testing.T) {
	{
		m := New("smtp.example.com", 587, "user", "pass", "from@example.com", "to@example.com")
		m.Subject("Test Subject")
		expected := "Test Subject"
		_assert.Equal(t, expected, m.subject)
	}
	{
		m := New("smtp.example.com", 587, "user", "pass", "from@example.com", "to@example.com")
		m.Subject("")
		expected := ""
		_assert.Equal(t, expected, m.subject)
	}
}
func TestContent(t *testing.T) {
	{
		m := New("smtp.example.com", 587, "user", "pass", "from@example.com", "to@example.com")
		m.Content("This is the body content")
		expected := "This is the body content"
		_assert.Equal(t, expected, m.content)
	}
	{
		m := New("smtp.example.com", 587, "user", "pass", "from@example.com", "to@example.com")
		m.Content("")
		expected := ""
		_assert.Equal(t, expected, m.content)
	}
}
func TestCc(t *testing.T) {
	{
		m := New("smtp.example.com", 587, "user", "pass", "from@example.com", "to@example.com")
		m.Cc("cc1@example.com")
		expected := []string{"cc1@example.com"}
		_assert.EqualByList(t, expected, m.ccList)
	}

	{
		m := New("smtp.example.com", 587, "user", "pass", "from@example.com", "to@example.com")
		m.Cc("cc1@example.com", "cc2@example.com")
		expected := []string{"cc1@example.com", "cc2@example.com"}
		_assert.EqualByList(t, expected, m.ccList)
	}

	{
		m := New("smtp.example.com", 587, "user", "pass", "from@example.com", "to@example.com")
		m.Cc("cc1@example.com", "cc1@example.com")
		expected := []string{"cc1@example.com", "cc1@example.com"}
		_assert.EqualByList(t, expected, m.ccList)
	}
}
func TestAttach(t *testing.T) {
	{
		attachmentPath := "test_attachment_1.txt"
		content := "This is a test attachment file"
		_file.Write(attachmentPath, content, 0644)
		defer _file.Delete(attachmentPath)
		m := New("smtp.example.com", 587, "user", "pass", "from@example.com", "to@example.com")
		m.Attach(attachmentPath)
		expected := []string{attachmentPath}
		_assert.EqualByList(t, expected, m.attachList)
	}
	{
		attachmentPath1 := "test_attachment_2_1.txt"
		content1 := "This is the first test attachment"
		_file.Write(attachmentPath1, content1, 0644)
		defer _file.Delete(attachmentPath1)
		attachmentPath2 := "test_attachment_2_2.txt"
		content2 := "This is the second test attachment"
		_file.Write(attachmentPath2, content2, 0644)
		defer _file.Delete(attachmentPath2)
		m := New("smtp.example.com", 587, "user", "pass", "from@example.com", "to@example.com")
		m.Attach(attachmentPath1, attachmentPath2)
		expected := []string{attachmentPath1, attachmentPath2}
		_assert.EqualByList(t, expected, m.attachList)
	}
	{
		attachmentPath := "test_attachment_3.txt"
		content := "This is a test attachment file"
		_file.Write(attachmentPath, content, 0644)
		defer _file.Delete(attachmentPath)
		m := New("smtp.example.com", 587, "user", "pass", "from@example.com", "to@example.com")
		m.Attach(attachmentPath)
		m.Attach(attachmentPath)
		expected := []string{attachmentPath, attachmentPath}
		_assert.EqualByList(t, expected, m.attachList)
	}
}
func TestSend(t *testing.T) {
	{
		New(
			"smtp.126.com",
			25,
			"guoguo_1234567890@126.com",
			"FZQJRCOHGWXSESMU",
			"guoguo_1234567890@126.com",
			"507979696@qq.com").
			Subject("subject").
			Content("<img src='https://www.baidu.com/img/flexible/logo/pc/result@2.png'/>").
			Send()
	}
}
