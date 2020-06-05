import axios from 'axios'

const sendEvCrawlab = async (eventCategory, eventAction, eventLabel) => {
}

export default {
  sendPv (page) {
    if (localStorage.getItem('useStats') !== '0') {
      // window._hmt.push(['_trackPageview', page])
      sendEvCrawlab('访问页面', page, '')
    }
  },
  sendEv (category, eventName, optLabel, optValue) {
    if (localStorage.getItem('useStats') !== '0') {
      // window._hmt.push(['_trackEvent', category, eventName, optLabel, optValue])
      sendEvCrawlab(category, eventName, optLabel)
    }
  }
}
