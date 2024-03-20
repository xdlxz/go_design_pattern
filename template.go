package main

import "fmt"

/*
模版方法模式使用继承机制，把通用步骤和通用方法放到父类中，把具体实现延迟到子类中实现。使得实现符合开闭原则(对扩展开放，对修改关闭)。
如实例代码中通用步骤在父类中实现（准备、下载、保存、收尾）下载和保存的具体实现留到子类中，并且提供 保存方法的默认实现。

因为子类要复用父类的通用方法，所以子类需要匿名组合父类。
因为父类要调用子类的具体方法，所以父类需要持有子类的引用（因为有多个子类，引用类型为interface）。
c++是父类隐式调用子类方法（多态），go是父类显式调用子类方法。

如果不用匿名组合呢？
*/

type Downloader interface {
	Download(uri string)
}

type template struct { // 模板类，父类
	implement // 引用子类
	uri       string
}

type implement interface {
	download()
	save()
}

func newTemplate(impl implement) *template {
	return &template{
		implement: impl,
	}
}

func (t *template) Download(uri string) {
	t.uri = uri
	fmt.Print("prepare downloading\n")
	t.implement.download()
	t.implement.save() // 显示调用子类方法，如果子类没实现，就调用父类默认方法
	fmt.Print("finish downloading\n")
}

func (t *template) save() {
	fmt.Print("default save\n")
}

////////////////具体实现类//////////////////////////
type HTTPDownloader struct {
	*template
	// tpl *template,如果不是匿名组合，就得显示通过成员变量的方式调用，无法使像使用自己的method一样使用父类的method
}

func NewHTTPDownloader() Downloader {
	downloader := &HTTPDownloader{}
	template := newTemplate(downloader)
	downloader.template = template
	//downloader.tpl = template
	return downloader
}

func (d *HTTPDownloader) download() {
	fmt.Printf("download %s via http\n", d.uri)
}

func (*HTTPDownloader) save() {
	fmt.Printf("http save\n")
}

type FTPDownloader struct {
	*template
}

func NewFTPDownloader() Downloader {
	downloader := &FTPDownloader{}
	template := newTemplate(downloader)
	downloader.template = template
	return downloader
}

func (d *FTPDownloader) download() {
	fmt.Printf("download %s via ftp\n", d.uri)
}

func main() {
	downloader := NewHTTPDownloader()
	downloader.Download("http://www.baidu.com")

	fmt.Println("///////////////////")
	downloader = NewFTPDownloader()
	downloader.Download("ftp://www.baidu.com")
}
