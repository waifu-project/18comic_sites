const got = require('got')
const cheerio = require('cheerio')

/**
 * 获取镜像
 */
const getMirror = async ()=> {
  try {
    const html_url = 'http://jmcomic.xyz'
    const res = await got(html_url)
    const data = res.body || ""
    const $ = cheerio.load(data)
    const listsEle = Array.from($('.has-inline-color.has-luminous-vivid-orange-color'))
    let mirrors = listsEle.map(item=> {
      const p = $(item).parent()
      let t = p.text().trim()
      const i = $(item).text().trim()
      if (i == t) {
        t = $(item).parent().parent().text().trim()
      }
      t = t.replace(i, "")
      // console.log('t: ', t)
      // console.log('i: ', i)
      return { title: t, api: i }
    }).filter(item=> {
      const flag = item.api.indexOf('18comic') != -1
      return flag
    })
    // console.log(mirrors)
    return mirrors
  } catch (e) {
    console.log('未知错误: ', e)
    return []
  }
}

module.exports = {
  getMirror
}