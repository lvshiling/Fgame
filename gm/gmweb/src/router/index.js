import Vue from 'vue'
import Router from 'vue-router'
Vue.config.devtools = true;
Vue.use(Router)

/* Layout */
import Layout from '@/views/layout/Layout'

/* Router Modules */
// import componentsRouter from './modules/components'
// import chartsRouter from './modules/charts'
// import tableRouter from './modules/table'
// import nestedRouter from './modules/nested'

/** note: Submenu only appear when children.length>=1
 *  detail see  https://panjiachen.github.io/vue-element-admin-site/guide/essentials/router-and-nav.html
 **/

/**
* hidden: true                   if `hidden:true` will not show in the sidebar(default is false)
* alwaysShow: true               if set true, will always show the root menu, whatever its child routes length
*                                if not set alwaysShow, only more than one route under the children
*                                it will becomes nested mode, otherwise not show the root menu
* redirect: noredirect           if `redirect:noredirect` will no redirect in the breadcrumb
* name:'router-name'             the name is used by <keep-alive> (must set!!!)
* meta : {
    roles: ['admin','editor']     will control the page roles (you can set multiple roles)
    title: 'title'               the name show in submenu and breadcrumb (recommend set)
    icon: 'svg-name'             the icon show in the sidebar,
    noCache: true                if true ,the page will no be cached(default is false)
  }
**/
export const constantRouterMap = [
  {
    path: '/redirect',
    component: Layout,
    hidden: true,
    children: [
      {
        path: '/redirect/:path*',
        component: () => import('@/views/redirect/index')
      }
    ]
  },
  {
    path: '/login',
    component: () => import('@/views/login/index'),
    hidden: true
  },
  {
    path: '/auth-redirect',
    component: () => import('@/views/login/authredirect'),
    hidden: true
  },
  {
    path: '/404',
    component: () => import('@/views/errorPage/404'),
    hidden: true
  },
  {
    path: '/401',
    component: () => import('@/views/errorPage/401'),
    hidden: true
  },
  {
    path: '',
    component: Layout,
    redirect: 'dashboard',
    children: [
      {
        path: 'dashboard',
        component: () => import('@/views/dashboard/index'),
        name: 'Dashboard',
        meta: { title: 'dashboard', icon: 'dashboard', noCache: true }
      }
    ]
  }
]

export default new Router({
  // mode: 'history', // require service support
  scrollBehavior: () => ({ y: 0 }),
  routes: constantRouterMap
})

export const asyncRouterMap = [
  {
    path: '/user',
    component: Layout,
    redirect: '/user/list',
    name: 'basemanage',
    alwaysShow: true,
    meta: { title: 'basemanage', icon: 'example', roles: ['super_admin', 'super_channel', 'channel', 'platform'] },
    children: [
      {
        path: 'user',
        name: 'usermanagebk',
        component: () => import('@/views/user/list'),
        meta: { title: 'usermanagebk', icon: '', roles: ['super_admin', 'super_channel', 'channel', 'platform'] }
      },
      {
        path: 'channel',
        name: 'channelmanage',
        component: () => import('@/views/channel/list'),
        meta: { title: 'channelmanage', icon: '', roles: ['super_admin'] }
      },
      {
        path: 'platform',
        name: 'platformmanage',
        component: () => import('@/views/platform/list'),
        meta: { title: 'platformmanage', icon: '', roles: ['super_admin'] }
      }
    ]
  },
  {
    path: '/gameplayer',
    component: Layout,
    redirect: '/gameplayer/playerlist',
    name: 'gameplayer',
    meta: { title: 'gameplayer', icon: 'example', roles: ['super_admin', 'super_channel', 'channel', 'platform', 'service', 'minitor','super_channel_service','common_service','gaoji_service','neigua','gs'] },
    children: [
      {
        path: 'player',
        name: 'playersearch',
        component: () => import('@/views/game/playerlist'),
        meta: { title: 'playersearch', icon: '', roles: ['super_admin', 'super_channel', 'channel', 'platform', 'service','minitor','super_channel_service','common_service','gaoji_service','neigua'] }
      },
      {
        path: 'playerInfo',
        name: 'playerInfo',
        hidden: true,
        component: () => import('@/views/game/player/playermain'),
        meta: { title: 'playersearch', icon: '', roles: ['super_admin', 'super_channel', 'channel', 'platform', 'service','minitor','super_channel_service','common_service','gaoji_service','neigua'] }
      },
      // {
      //   path: 'monitor',
      //   name: '测试监控',
      //   component: () => import('@/views/monitor/monitor'),
      //   meta: { title: '测试监控', icon: 'table' }
      // },
      {
        path: 'chat',
        name: 'chat',
        component: () => import('@/views/monitor/chat'),
        meta: { title: 'chat', icon: '', roles: ['super_admin', 'super_channel', 'channel', 'platform', 'service', 'minitor','super_channel_service','common_service','gaoji_service','gs'] }
      },
      {
        path: 'chatset',
        name: 'chatset',
        component: () => import('@/views/monitor/chatset'),
        meta: { title: 'chatset', icon: '', roles: ['super_admin', 'super_channel', 'channel', 'platform', 'service','super_channel_service','common_service','gaoji_service'] }
      },
      {
        path: 'alliance',
        name: 'alliance',
        component: () => import('@/views/game/alliance'),
        meta: { title: 'alliance', icon: '', roles: ['super_admin', 'super_channel', 'channel', 'platform', 'service','minitor','super_channel_service','common_service','gaoji_service'] }
      },
      {
        path: 'applymail',
        name: 'applymail',
        component: () => import('@/views/manage/applymail'),
        meta: { title: 'applymail', icon: '', roles: ['super_admin', 'super_channel', 'channel', 'platform', 'service','super_channel_service','common_service','gaoji_service'] }
      },
      {
        path: 'approvemail',
        name: 'approvemail',
        component: () => import('@/views/manage/approvemail'),
        meta: { title: 'approvemail', icon: '', roles: ['super_admin', 'super_channel', 'channel', 'platform', 'service','super_channel_service','common_service','gaoji_service'] }
      },
      {
        path: 'serversupportpool',
        name: 'serversupportpool',
        component: () => import('@/views/manage/serversupportpool'),
        meta: { title: 'serversupportpool', icon: '', roles: ['super_admin', 'super_channel', 'channel', 'platform', 'service'] }
      },
      {
        path: 'serversupportplayer',
        name: 'serversupportplayer',
        component: () => import('@/views/manage/serversupportplayer'),
        meta: { title: 'serversupportplayer', icon: '', roles: ['super_admin', 'super_channel', 'channel', 'platform', 'service','minitor','neigua'] }
      },
      {
        path: 'serverlogin',
        name: 'serverlogin',
        component: () => import('@/views/manage/serverlogin'),
        meta: { title: 'serverlogin', icon: '', roles: ['super_admin', 'super_channel', 'channel', 'platform', 'service','minitor'] }
      },
      {
        path: 'doublecharge',
        name: 'doublecharge',
        component: () => import('@/views/manage/doublecharge'),
        meta: { title: 'doublecharge', icon: '', roles: ['super_admin'] }
      },
      {
        path: 'orderlist',
        name: 'orderlist',
        component: () => import('@/views/center/order'),
        meta: { title: 'orderlist', icon: '', roles: ['super_admin', 'super_channel', 'channel', 'platform', 'service','common_service','gaoji_service'] }
      },
      {
        path: 'orderstatic',
        name: 'orderstatic',
        component: () => import('@/views/manage/orderstatic'),
        meta: { title: 'orderstatic', icon: '', roles: ['super_admin', 'super_channel', 'channel', 'platform', 'service'] }
      },
      {
        path: 'ordersignlequery',
        name: 'ordersignlequery',
        component: () => import('@/views/manage/ordestatic/orderlist'),
        meta: { title: 'ordersignlequery', icon: '', roles: ['super_admin', 'super_channel', 'channel', 'platform','gaoji_service'] }
      },
      {
        path: 'onlinereport',
        name: 'onlinereport',
        component: () => import('@/views/report/online'),
        meta: { title: 'onlinereport', icon: '', roles: ['super_admin', 'super_channel', 'channel', 'platform', 'service','common_service','gaoji_service'] }
      },
      {
        path: 'ngonlinereport',
        name: 'ngonlinereport',
        component: () => import('@/views/report/ngonline'),
        meta: { title: 'ngonlinereport', icon: '', roles: ['super_admin', 'super_channel', 'channel', 'platform', 'service','common_service','gaoji_service'] }
      },
      {
        path: 'retentionreport',
        name: 'retentionreport',
        component: () => import('@/views/report/retention'),
        meta: { title: 'retentionreport', icon: '', roles: ['super_admin'] }
      },
      {
        path: 'notice',
        name: 'notice',
        component: () => import('@/views/manage/notice'),
        meta: { title: 'notice', icon: '', roles: ['super_admin', 'super_channel', 'channel', 'platform', 'service','super_channel_service'] }
      },
      {
        path: 'redeem',
        name: 'redeem',
        component: () => import('@/views/manage/redeem'),
        meta: { title: 'redeem', icon: '', roles: ['super_admin', 'super_channel', 'channel', 'platform', 'super_channel_service'] }
      },
      {
        path: 'centeruser',
        name: 'centeruser',
        component: () => import('@/views/center/userquery'),
        meta: { title: 'centeruser', icon: '', roles: ['super_admin', 'super_channel', 'channel', 'platform', 'service','super_channel_service','common_service','gaoji_service'] }
      },
      {
        path: 'centernewguauser',
        name: 'centernewguauser',
        component: () => import('@/views/center/userneigua'),
        meta: { title: 'centernewguauser', icon: '', roles: ['super_admin', 'super_channel', 'channel', 'platform', 'service','super_channel_service','common_service','gaoji_service','neigua'] }
      },
      // {
      //   path: 'goldchange',
      //   name: 'goldchange',
      //   component: () => import('@/views/manage/goldchange'),
      //   meta: { title: 'goldchange', icon: '', roles: ['super_admin'] }
      // },
      // {
      //   path: 'goldreduce',
      //   name: 'goldreduce',
      //   component: () => import('@/views/manage/goldreduce'),
      //   meta: { title: 'goldreduce', icon: '', roles: ['super_admin'] }
      // },
      {
        path: 'newbindgold',
        name: 'newbindgold',
        component: () => import('@/views/manage/newbindgold'),
        meta: { title: 'newbindgold', icon: '', roles: ['super_admin'] }
      },
      {
        path: 'newbindgoldreduce',
        name: 'newbindgoldreduce',
        component: () => import('@/views/manage/newbindgoldreduce'),
        meta: { title: 'newbindgoldreduce', icon: '', roles: ['super_admin'] }
      },
      {
        path: 'newgoldchange',
        name: 'newgoldchange',
        component: () => import('@/views/manage/newgoldchange'),
        meta: { title: 'newgoldchange', icon: '', roles: ['super_admin'] }
      },
      {
        path: 'newgoldchangereduce',
        name: 'newgoldchangereduce',
        component: () => import('@/views/manage/newgoldchangereduce'),
        meta: { title: 'newgoldchangereduce', icon: '', roles: ['super_admin'] }
      },
      {
        path: 'redeemcode/:id(\\d+)',
        name: 'redeemcode',
        hidden: true,
        component: () => import('@/views/manage/redeemcode'),
        meta: { title: 'redeemcode', icon: '', roles: ['super_admin', 'super_channel', 'channel', 'platform', 'service','common_service','gaoji_service','super_channel_service'] }
      },
      {
        path: 'centerserverquery',
        name: 'centerserverquery',
        component: () => import('@/views/center/serverquery'),
        meta: { title: 'centerserverquery', icon: '', roles: ['super_admin', 'super_channel', 'channel', 'platform', 'service','super_channel_service'] }
      },
      {
        path: 'serverplayerlevel',
        name: 'serverplayerlevel',
        component: () => import('@/views/game/playerlevel'),
        meta: { title: 'serverplayerlevel', icon: '', roles: ['super_admin'] }
      },
      {
        path: 'playerquest',
        name: 'playerquest',
        component: () => import('@/views/game/playerquest'),
        meta: { title: 'playerquest', icon: '', roles: ['super_admin'] }
      },
      {
        path: 'tradequery',
        name: 'tradequery',
        component: () => import('@/views/center/tradeitem'),
        meta: { title: 'tradequery', icon: '', roles: ['super_admin'] }
      },
      {
        path: 'serverdaily',
        name: 'serverdaily',
        component: () => import('@/views/manage/dailyserverreport'),
        meta: { title: 'serverdaily', icon: '', roles: ['super_admin', 'super_channel', 'channel', 'platform', 'service','super_channel_service','common_service','gaoji_service'] }
      },
      {
        path: 'zhanquServer',
        name: 'zhanquServer',
        component: () => import('@/views/center/server_zhanqu'),
        meta: { title: 'zhanquServer', icon: '', roles: ['super_admin', 'super_channel', 'channel', 'platform', 'service','super_channel_service','common_service','gaoji_service'] }
      },
      {
        path: 'ipquery',
        name: 'ipquery',
        component: () => import('@/views/center/ipquery'),
        meta: { title: 'ipquery', icon: '', roles: ['super_admin', 'super_channel', 'channel', 'platform', 'service','super_channel_service','common_service','gaoji_service','minitor'] }
      },
      {
        path: 'recycle',
        name: 'recycle',
        component: () => import('@/views/game/recycle'),
        meta: { title: 'recycle', icon: '', roles: ['super_admin', 'super_channel', 'channel', 'platform', 'service','super_channel_service','common_service','gaoji_service'] }
      },
      {
        path: 'platformSupportPoolSet',
        name: 'platformSupportPoolSet',
        component: () => import('@/views/manage/platformsupportpoolset'),
        meta: { title: 'platformSupportPoolSet', icon: '', roles: ['super_admin', 'super_channel', 'channel', 'platform', 'service'] }
      },
      {
        path: 'platformChatSet',
        name: 'platformChatSet',
        component: () => import('@/views/monitor/chatsetplatform'),
        meta: { title: 'platformChatSet', icon: '', roles: ['super_admin', 'super_channel', 'channel', 'platform', 'service'] }
      },
      {
        path: 'feedBackFee',
        name: 'feedBackFee',
        component: () => import('@/views/feedbackfee/feedbackfee'),
        meta: { title: 'feedBackFee', icon: '', roles: ['super_admin', 'super_channel', 'channel', 'platform', 'service'] }
      }
    ]
  },
  {
    path: '/centerset',
    component: Layout,
    redirect: '/centerset/platform',
    name: 'centerset',
    meta: { title: 'centerset', icon: 'example', roles: ['super_admin'] },
    children: [
      {
        path: 'centerplat',
        name: 'centerplat',
        component: () => import('@/views/center/platform'),
        meta: { title: 'centerplat', icon: '', roles: ['super_admin'] }
      },
      {
        path: 'server',
        name: 'server',
        component: () => import('@/views/center/server'),
        meta: { title: 'server', icon: '', roles: ['super_admin'] }
      },
      {
        path: 'serverset',
        name: 'serverset',
        hidden: true,
        component: () => import('@/views/center/serverset'),
        meta: { title: 'serverset', icon: '', roles: ['super_admin'] }
      },
      {
        path: 'usermanage',
        name: 'usermanage',
        component: () => import('@/views/center/user'),
        meta: { title: 'usermanage', icon: '', roles: ['super_admin'] }
      },
      {
        path: 'centernotice',
        name: 'centernotice',
        component: () => import('@/views/center/notice'),
        meta: { title: 'centernotice', icon: '', roles: ['super_admin'] }
      },
      {
        path: 'servergm',
        name: 'servergm',
        component: () => import('@/views/center/gmcmd'),
        meta: { title: 'servergm', icon: '', roles: ['super_admin'] }
      },
      {
        path: 'platformmarry',
        name: 'platformmarry',
        component: () => import('@/views/center/platformmarry'),
        meta: { title: 'platformmarry', icon: '', roles: ['super_admin'] }
      },
      {
        path: 'platformsetting',
        name: 'platformsetting',
        component: () => import('@/views/center/platformset'),
        meta: { title: 'platformsetting', icon: '', roles: ['super_admin'] }
      },
      {
        path: 'clientversionset',
        name: 'clientversionset',
        component: () => import('@/views/center/client_version'),
        meta: { title: 'clientversionset', icon: '', roles: ['super_admin'] }
      },
      {
        path: 'platformserverconfig',
        name: 'platformserverconfig',
        component: () => import('@/views/center/platform_server_config'),
        meta: { title: 'platformserverconfig', icon: '', roles: ['super_admin'] }
      },
      {
        path: 'jiaoYiZhanQu',
        name: 'jiaoYiZhanQu',
        component: () => import('@/views/center/jiaoyizhanqu'),
        meta: { title: 'jiaoYiZhanQu', icon: '', roles: ['super_admin'] }
      },
      {
        path: 'jiaoYiZhanQuServerSet',
        name: 'jiaoYiZhanQuServerSet',
        component: () => import('@/views/center/serverjiaoyizhanqu'),
        meta: { title: 'jiaoYiZhanQuServerSet', icon: '', roles: ['super_admin'] }
      },
      {
        path: 'quanPingTaiServerSet',
        name: 'quanPingTaiServerSet',
        component: () => import('@/views/center/serverpingtaifu'),
        meta: { title: 'quanPingTaiServerSet', icon: '', roles: ['super_admin'] }
      },
      {
        path: 'zhanQuServerSet',
        name: 'zhanQuServerSet',
        component: () => import('@/views/center/serverzhanqu'),
        meta: { title: 'zhanQuServerSet', icon: '', roles: ['super_admin'] }
      },
      {
        path: 'chengZhanServerSet',
        name: 'chengZhanServerSet',
        component: () => import('@/views/center/serverchengzhan'),
        meta: { title: 'chengZhanServerSet', icon: '', roles: ['super_admin'] }
      }
    ]
  },
  {
    path: '/log',
    component: Layout,
    redirect: '/log/list',
    name: 'log',
    meta: { title: 'log', icon: 'example', roles: ['super_admin', 'super_channel', 'channel', 'platform', 'service','super_channel_service','common_service','gaoji_service'] },
    children: [
      {
        path: 'playerlog',
        name: 'playerlog',
        component: () => import('@/views/log/playerlist'),
        meta: { title: 'playerlog', icon: '', roles: ['super_admin', 'super_channel', 'channel', 'platform', 'service','super_channel_service','common_service','gaoji_service'] }
      },
      {
        path: 'serverlog',
        name: 'serverlog',
        component: () => import('@/views/log/serverlist'),
        meta: { title: 'serverlog', icon: '', roles: ['super_admin', 'super_channel', 'channel', 'platform', 'service','super_channel_service'] }
      },
      {
        path: 'alliancelog',
        name: 'alliancelog',
        component: () => import('@/views/log/alliancelist'),
        meta: { title: 'alliancelog', icon: '', roles: ['super_admin', 'super_channel', 'channel', 'platform', 'service','super_channel_service','common_service','gaoji_service'] }
      },
      {
        path: 'tradelog',
        name: 'tradelog',
        component: () => import('@/views/log/tradelist'),
        meta: { title: 'tradelog', icon: '', roles: ['super_admin', 'super_channel', 'channel', 'platform', 'service','super_channel_service','common_service','gaoji_service'] }
      },
      {
        path: 'playerstats',
        name: 'playerstats',
        component: () => import('@/views/log/playerstats'),
        meta: { title: 'playerstats', icon: '', roles: ['super_admin'] }
      }
    ]
  },
  { path: '*', redirect: '/404', hidden: true }
]
