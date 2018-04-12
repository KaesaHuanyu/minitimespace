//app.js
App({
  onLaunch: function () {
    var session_key = wx.getStorageSync('session_key')
    if (session_key == '') {
      // 登录
      wx.login({
        success: res => {
          // 发送 res.code 到后台换取 openId, sessionKey, unionId
          if (res.code) {
            wx.request({
              url: this.globalData.host + '/login?code=' + res.code,
              success: resp => {
                wx.setStorageSync('session_key', resp.data.data.session_key)
              }
            })
          } else {
            console.log("login failed")
          }
        }
      })
    }

    // 获取用户信息
    wx.getSetting({
      success: res => {
        if (res.authSetting['scope.userInfo']) {
          // 已经授权，可以直接调用 getUserInfo 获取头像昵称，不会弹框
          wx.getUserInfo({
            lang: 'zh_CN',
            success: res => {
              // 可以将 res 发送给后台解码出 unionId
              this.globalData.userInfo = res.userInfo
              wx.request({
                url: this.globalData.host + '/users?session_key=' + session_key,
                method: 'POST',
                json: true,
                data: res.userInfo,
                success: resp => {
                  if (resp.data.code == 6) {
                    // 登录
                    wx.login({
                      success: res => {
                        // 发送 res.code 到后台换取 openId, sessionKey, unionId
                        if (res.code) {
                          wx.request({
                            url: this.globalData.host + '/login?code=' + res.code,
                            success: resp => {
                              wx.setStorageSync('session_key', resp.data.data.session_key)
                              wx.request({
                                url: this.globalData.host + '/users?session_key=' + session_key,
                                method: 'POST',
                                json: true,
                                data: res.userInfo
                              })
                            }
                          })
                        } else {
                          console.log("login failed")
                        }
                      }
                    })
                  }
                }
              })

              // 由于 getUserInfo 是网络请求，可能会在 Page.onLoad 之后才返回
              // 所以此处加入 callback 以防止这种情况
              if (this.userInfoReadyCallback) {
                this.userInfoReadyCallback(res)
              }
            }
          })
        }
      }
    })
  },
  globalData: {
    userInfo: null,
    host: 'http://localhost:8823'
  }
})