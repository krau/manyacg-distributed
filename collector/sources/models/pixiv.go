package models

import "encoding/xml"



type PixivRss struct {
	XMLName xml.Name `xml:"rss"`
	Channel channel  `xml:"channel"`
}

type channel struct {
	XMLName xml.Name `xml:"channel"`
	Items   []Item   `xml:"item"`
}

type Item struct {
	XMLName     xml.Name `xml:"item"`
	Title       string   `xml:"title"`
	Description string   `xml:"description"`
	PubDate     string   `xml:"pubDate"`
	Guid        string   `xml:"guid"`
	Link        string   `xml:"link"`
	Author      string   `xml:"author"`
}


type PixivAjaxResp struct {
	Err     bool               `json:"error"`
	Message string             `json:"message"`
	Body    *PixivAjaxRespBody `json:"body"`
}

type PixivAjaxRespBody struct {
	IllustId   string                `json:"illustId"`
	IllustType int                   `json:"illustType"`
	Tags       PixivAjaxRespBodyTags `json:"tags"`
	UserId     string                `json:"userId"`
	ExtraData  PixivAjaxRespBodyExtraData `json:"extraData"`
}

type PixivAjaxRespBodyTags struct {
	Tags []PixivAjaxRespBodyTagsTag `json:"tags"`
}

type PixivAjaxRespBodyTagsTag struct {
	// 返回里确实就是这么套的
	Tag         string                           `json:"tag"`
	Translation *PixivAjaxRespBodyTagTranslation `json:"translation"`
}

type PixivAjaxRespBodyTagTranslation struct {
	// en翻译实际上是中文
	En string `json:"en"`
}


type PixivAjaxRespBodyExtraData struct {
	Meta PixivAjaxRespBodyExtraDataMeta `json:"meta"`
}

type PixivAjaxRespBodyExtraDataMeta struct {
	Description string `json:"description"`
}