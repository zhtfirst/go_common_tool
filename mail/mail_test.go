/*
 * @Author: gavin v_zhangtao15@tal.com
 * @Date: 2025-04-24 17:00:30
 * @LastEditors: gavin v_zhangtao15@tal.com
 * @LastEditTime: 2025-05-19 16:55:30
 * @FilePath: /mail/mail_test.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package mail

import (
	"testing"
)

func TestSend(t *testing.T) {
	// 设置测试邮件内容
	subject := "Test Email"
	body := "This is a test email"
	recipient := "recipient@example.com"

	// 创建一个自定义的Mail结构体，覆盖Send方法的行为
	// 在实际项目中，可以使用mock库来实现更完整的模拟
	customMail := &Mail{
		Host:     "smtp.example.com",
		Port:     465,
		Username: "test@example.com",
		Password: "password",
	}

	// 添加附件和嵌入图片
	err := customMail.Send(recipient, subject, body,
		SetAttachment("test.docx"),
		SetAttachment("test.pdf", SetAttachmentName("测试.pdf")),
		SetEmbedded("test.png", SetEmbeddedName("测试图片.png")))

	// 验证基本参数
	if customMail.Host != "smtp.example.com" {
		t.Errorf("Expected host smtp.example.com, got %s", customMail.Host)
	}
	if customMail.Port != 465 {
		t.Errorf("Expected port 465, got %d", customMail.Port)
	}

	// 验证附件是否正确添加
	if len(customMail.attachments) != 2 {
		t.Errorf("Expected 2 attachments, got %d", len(customMail.attachments))
	}

	// 验证嵌入图片是否正确添加
	if len(customMail.embedded) != 1 {
		t.Errorf("Expected 1 embedded file, got %d", len(customMail.embedded))
	}

	// 验证附件名称
	if customMail.attachments[0].file != "test.docx" {
		t.Errorf("Expected attachment file test.docx, got %s", customMail.attachments[0].file)
	}

	if customMail.attachments[1].file != "test.pdf" || customMail.attachments[1].name != "测试.pdf" {
		t.Errorf("Expected attachment test.pdf with name 测试.pdf, got file %s with name %s",
			customMail.attachments[1].file, customMail.attachments[1].name)
	}

	// 验证嵌入图片
	if customMail.embedded[0].file != "test.png" || customMail.embedded[0].name != "测试图片.png" {
		t.Errorf("Expected embedded test.png with name 测试图片.png, got file %s with name %s",
			customMail.embedded[0].file, customMail.embedded[0].name)
	}

	// 在实际项目中，这里应该有错误，因为我们没有真正发送邮件
	// 但在测试环境中，我们只关心参数设置是否正确
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}
