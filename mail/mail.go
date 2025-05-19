package mail

import (
	"strings"

	"gopkg.in/gomail.v2"
)

// Mail represents an email configuration
type Mail struct {
	Host        string
	Port        int
	Username    string
	Password    string
	attachments []*MailFile
	embedded    []*MailFile
	cc          []string
	bcc         []string
}

// MailFile represents a file attachment or embedded file
type MailFile struct {
	file string
	name string
}

// MailOption defines a function type for mail options
type MailOption func(*Mail)

// FileSetting defines a function type for file settings
type FileSetting func(*MailFile)

// NewMail creates a new Mail instance with the provided configuration
func NewMail(host string, port int, username string, password string) *Mail {
	return &Mail{
		Host:     host,
		Port:     port,
		Username: username,
		Password: password,
	}
}

// Send 发送邮件，支持指定收件人、主题、内容和其他可选项
// 参数:
//   - to: 收件人邮箱地址，多个地址用逗号分隔
//   - subject: 邮件主题
//   - body: 邮件正文内容（HTML格式）
//   - opts: 可选配置项，如附件、抄送、密送等
//
// 返回:
//   - error: 发送成功返回nil，失败返回相应错误
func (m *Mail) Send(to string, subject string, body string, opts ...MailOption) error {
	gm := gomail.NewMessage()

	// Apply all options
	for _, opt := range opts {
		opt(m)
	}

	// Set sender
	gm.SetHeader("From", m.Username)

	// Set recipients
	mailArrTo := strings.Split(to, ",")
	gm.SetHeader("To", mailArrTo...)

	// Set subject
	gm.SetHeader("Subject", subject)

	// Set email body
	gm.SetBody("text/html", body)

	// Set CC recipients
	if len(m.cc) > 0 {
		gm.SetHeader("Cc", m.cc...)
	}

	// Set BCC recipients
	if len(m.bcc) > 0 {
		gm.SetHeader("Bcc", m.bcc...)
	}

	// Add attachments
	for _, attachment := range m.attachments {
		// 如果attachment.name为空,则不调用rename todo优化
		if attachment.name == "" {
			gm.Attach(attachment.file)
		} else {
			gm.Attach(attachment.file, gomail.Rename(attachment.name))
		}
	}

	// Add embedded files
	for _, embedded := range m.embedded {
		// 如果embedded.name为空,则不调用rename todo优化
		if embedded.name == "" {
			gm.Embed(embedded.file)
		} else {
			gm.Embed(embedded.file, gomail.Rename(embedded.name))
		}
	}

	// Create dialer and send email
	d := gomail.NewDialer(m.Host, m.Port, m.Username, m.Password)

	return d.DialAndSend(gm)
}

// SetAttachment adds an attachment to the email
func SetAttachment(filename string, fileSetting ...FileSetting) MailOption {
	return func(m *Mail) {
		mf := &MailFile{
			file: filename,
		}

		for _, setting := range fileSetting {
			setting(mf)
		}

		m.attachments = append(m.attachments, mf)
	}
}

// SetEmbedded adds an embedded file to the email
func SetEmbedded(filename string, fileSetting ...FileSetting) MailOption {
	return func(m *Mail) {
		mf := &MailFile{
			file: filename,
		}

		for _, setting := range fileSetting {
			setting(mf)
		}

		m.embedded = append(m.embedded, mf)
	}
}

// SetCC adds a CC recipient to the email
func SetCC(cc string) MailOption {
	return func(m *Mail) {
		m.cc = append(m.cc, cc)
	}
}

// SetBcc adds a BCC recipient to the email
func SetBcc(bcc string) MailOption {
	return func(m *Mail) {
		m.bcc = append(m.bcc, bcc)
	}
}

// SetAttachmentName sets a name for an attachment
func SetAttachmentName(name string) FileSetting {
	return func(mf *MailFile) {
		mf.name = name
	}
}

// SetEmbeddedName sets a name for an embedded file
func SetEmbeddedName(name string) FileSetting {
	return func(mf *MailFile) {
		mf.name = name
	}
}
