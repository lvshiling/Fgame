<template>
  <div class="app-container">
    <el-tabs type="border-card">
      <el-tab-pane label="聊天监控">
        <div>
          <div class="filter-container">
            <el-select v-model="chatListQuery.platformId" placeholder="平台" style="width: 160px" class="filter-item" @change="handlePlatformChange">
              <el-option v-for="item in platformList" :key="item.platformId" :label="item.platformName" :value="item.platformId" />
            </el-select>
            <el-select v-model="chatListQuery.serverArray" multiple collapse-tags placeholder="服务器" clearable style="width: 220px" class="filter-item" @change="handleServerChange">
              <el-option v-for="item in serverList" :key="item.id" :label="item.serverName" :value="item.id"/>
            </el-select>
            <el-input v-model.trim="chatListQuery.playerName" placeholder="角色名" style="width: 200px;" class="filter-item" />
            <el-input v-model.trim="chatListQuery.toPlayerName" placeholder="私聊目标" style="width: 200px;" class="filter-item" />
            <el-input v-model="chatListQuery.vipLevel" placeholder="VIP等级" style="width: 200px;" class="filter-item" />

            <el-select v-model="chatListQuery.chatType" placeholder="频道" clearable style="width: 120px" class="filter-item">
              <el-option v-for="item in chatTypeList" :key="item.key" :label="item.name" :value="item.key"/>
            </el-select>
            <el-input v-model.trim="chatListQuery.chatMsg" placeholder="聊天内容" style="width: 200px;" class="filter-item" />

            <el-button v-waves :type="chatListQuery.socketbtnType" class="filter-item" icon="el-icon-search" @click="handleFilter">{{ chatListQuery.socketbtnMsg }}</el-button>
            <el-button v-waves type="primary" class="filter-item" icon="el-icon-search" @click="handleSearch">刷新</el-button>
            <el-button class="filter-item" style="margin-left: 10px;" type="primary" icon="el-icon-edit" @click="handleMinGan">敏感词</el-button>
          </div>
        </div>
        <el-table
          :key="chatKey"
          :data="chatList"
          border
          fit
          highlight-current-row
          style="width: 100%;margin-top:15px;">
          <el-table-column label="玩家Id" align="center" width="180px">
            <template slot-scope="scope">
              <span>{{ scope.row.playerId }}</span>
            </template>
          </el-table-column>
          <el-table-column label="角色名" align="center" width="180px">
            <template slot-scope="scope">
              <span>{{ scope.row.playerName }}</span>
            </template>
          </el-table-column>
          <el-table-column label="服务器Id" width="100px" align="center">
            <template slot-scope="scope">
              <span>{{ scope.row.centerServerId }}</span>
            </template>
          </el-table-column>
          <el-table-column label="ip" width="120px" align="center">
            <template slot-scope="scope">
              <span>{{ scope.row.ip }}</span>
            </template>
          </el-table-column>
          <el-table-column label="VIP等级" width="100px" align="center">
            <template slot-scope="scope">
              <span>{{ scope.row.vipLevel }}</span>
            </template>
          </el-table-column>
          <el-table-column label="等级" width="60px">
            <template slot-scope="scope">
              <span>{{ scope.row.playerLevel }}</span>
            </template>
          </el-table-column>
          <el-table-column label="频道" width="60px">
            <template slot-scope="scope">
              <span>{{ scope.row.chatType | chatTypeFilter }}</span>
            </template>
          </el-table-column>
          <el-table-column label="私聊目标" width="180px">
            <template slot-scope="scope">
              <span>{{ scope.row.toPlayerName }}</span>
            </template>
          </el-table-column>
          <el-table-column label="聊天方式" width="100px">
            <template slot-scope="scope">
              <span>{{ scope.row.chatMethod | chatMethodFilter }}</span>
            </template>
          </el-table-column>
          <el-table-column label="聊天内容" min-width="150px">
            <template slot-scope="scope">
              <span v-html="scope.row.chatMsg"/>
            </template>
          </el-table-column>
          <el-table-column label="时间" width="160px">
            <template slot-scope="scope">
              <span>{{ scope.row.chatTime | parseTime }}</span>
            </template>
          </el-table-column>
          <el-table-column label="操作" align="center" width="500" class-name="small-padding fixed-width">
            <template slot-scope="scope">
              <el-button type="primary" size="mini" @click="handleForbid(scope.row)"
              v-permission="['super_admin', 'super_channel', 'channel', 'platform', 'service', 'minitor','super_channel_service','common_service','gaoji_service']">
              封禁</el-button>
              <el-button size="mini" type="primary" @click="handleForbidChat(scope.row)"
              v-permission="['super_admin', 'super_channel', 'channel', 'platform', 'service', 'minitor','super_channel_service','common_service','gaoji_service']">
              禁言</el-button>
              <el-button size="mini" type="primary" @click="handleIgnoreChat(scope.row)"
              v-permission="['super_admin', 'super_channel', 'channel', 'platform', 'service', 'minitor','super_channel_service','common_service','gaoji_service']">
              禁默</el-button>
              <el-button type="primary" size="mini" @click="handleKickOut(scope.row)"
              v-permission="['super_admin', 'super_channel', 'channel', 'platform', 'service', 'minitor','super_channel_service','common_service','gaoji_service']">
              踢人</el-button>
              <el-button size="mini" type="primary" @click="handleCenterForbid(scope.row)"
              v-permission="['super_admin', 'super_channel', 'channel', 'platform', 'service', 'minitor','super_channel_service','common_service','gaoji_service']">
              封号</el-button>
              <el-button size="mini" type="primary" @click="handleIpForbid(scope.row)"
              v-permission="['super_admin', 'super_channel', 'channel', 'platform', 'service', 'minitor','super_channel_service','common_service','gaoji_service']">
              禁ip</el-button>
            </template>
          </el-table-column>
        </el-table>

        <el-dialog :visible.sync="dialogFormVisible" title="敏感词">
          <!-- <el-form ref="dataForm" :model="sensitive" label-position="left" label-width="0px" style="width: 400px; margin-left:50px;">
                    <el-form-item>

                    </el-form-item>
                </el-form> -->
          <div style="height:400px;">
            <el-input v-model="sensitive.content" style="margin-bottom：10px;" minlength="400px" autosize type="textarea" placeholder="请输入敏感词，以英文逗号隔开"/>
          </div>
          <div slot="footer" class="dialog-footer">
            <el-button @click="dialogFormVisible = false">取消</el-button>
            <el-button type="primary" @click="updateSensitive">确定</el-button>
          </div>
        </el-dialog>

        <el-dialog :visible.sync="dialogForbidFormVisible" title="封禁用户">
          <el-form ref="dataForm" :model="monitorTemp" label-position="left" label-width="120px" style="width: 400px; margin-left:50px;">
            <el-form-item label="封禁用户名">
              <el-input v-model="monitorTemp.playerName" :disabled="true"/>
            </el-form-item>
            <el-form-item label="封禁原因">
              <el-input v-model="monitorTemp.reason"/>
            </el-form-item>
            <el-form-item label="封禁时长">
              <el-select v-model="monitorTemp.forbidTime" placeholder="封禁时长" style="width: 120px" class="filter-item">
                <el-option v-for="item in chatForbidTimeArray" :key="item.key" :label="item.name" :value="item.key"/>
              </el-select>
            </el-form-item>
            
          </el-form>
          <div slot="footer" class="dialog-footer">
            <el-button @click="dialogForbidFormVisible = false">取消</el-button>
            <el-button type="primary" @click="updateForbidPlayer">确定</el-button>
          </div>
        </el-dialog>

        <el-dialog :visible.sync="dialogForbidChatFormVisible" title="禁言用户">
          <el-form ref="dataForm" :model="monitorTemp" label-position="left" label-width="120px" style="width: 400px; margin-left:50px;">
            <el-form-item label="封禁用户名">
              <el-input v-model="monitorTemp.playerName" :disabled="true"/>
            </el-form-item>
            <el-form-item label="封禁原因">
              <el-input v-model="monitorTemp.reason"/>
            </el-form-item>
            <el-form-item label="封禁时长">
              <el-select v-model="monitorTemp.forbidTime" placeholder="封禁时长" style="width: 120px" class="filter-item">
                <el-option v-for="item in chatForbidTimeArray" :key="item.key" :label="item.name" :value="item.key"/>
              </el-select>
            </el-form-item>
          </el-form>
          <div slot="footer" class="dialog-footer">
            <el-button @click="dialogForbidChatFormVisible = false">取消</el-button>
            <el-button type="primary" @click="updateForbidChatPlayer">确定</el-button>
          </div>
        </el-dialog>

        <el-dialog :visible.sync="dialogIgnoreChatFormVisible" title="禁默用户">
          <el-form ref="dataForm" :model="monitorTemp" label-position="left" label-width="120px" style="width: 400px; margin-left:50px;">
            <el-form-item label="禁默用户名">
              <el-input v-model="monitorTemp.playerName" :disabled="true"/>
            </el-form-item>
            <el-form-item label="禁默原因">
              <el-input v-model="monitorTemp.reason"/>
            </el-form-item>
            <el-form-item label="禁默时长">
              <el-select v-model="monitorTemp.forbidTime" placeholder="禁默时长" style="width: 120px" class="filter-item">
                <el-option v-for="item in chatForbidTimeArray" :key="item.key" :label="item.name" :value="item.key"/>
              </el-select>
            </el-form-item>
          </el-form>
          <div slot="footer" class="dialog-footer">
            <el-button @click="dialogIgnoreChatFormVisible = false">取消</el-button>
            <el-button type="primary" @click="updateIgnoreChatPlayer">确定</el-button>
          </div>
        </el-dialog>

        <el-dialog :visible.sync="dialogKickVisible" title="踢人">
          <el-form ref="dataForm" :model="monitorTemp" label-position="left" label-width="120px" style="width: 400px; margin-left:50px;">
            <el-form-item label="用户名">
              <el-input v-model="monitorTemp.playerName" :disabled="true"/>
            </el-form-item>
            <el-form-item label="踢人原因">
              <el-input v-model="monitorTemp.reason"/>
            </el-form-item>
          </el-form>
          <div slot="footer" class="dialog-footer">
            <el-button @click="dialogKickVisible = false">取消</el-button>
            <el-button type="primary" @click="updateKickOutPlayer">确定</el-button>
          </div>
        </el-dialog>

        <el-dialog :visible.sync="dialogCenterForbidVisible" title="封号">
          <el-form ref="dataForm" :model="monitorTemp" label-position="left" label-width="120px" style="width: 400px; margin-left:50px;">
            <el-form-item label="封号用户名">
              <el-input v-model="monitorTemp.playerName" :disabled="true"/>
            </el-form-item>
            <el-form-item label="封号原因">
              <el-input v-model="monitorTemp.reason"/>
            </el-form-item>
            <el-form-item label="封号时长">
              <el-select v-model="monitorTemp.forbidTime" placeholder="封号时长" style="width: 120px" class="filter-item">
                <el-option v-for="item in chatForbidTimeArray" :key="item.key" :label="item.name" :value="item.key"/>
              </el-select>
            </el-form-item>
          </el-form>
          <div slot="footer" class="dialog-footer">
            <el-button @click="dialogCenterForbidVisible = false">取消</el-button>
            <el-button type="primary" @click="updateCenterUserForbid">确定</el-button>
          </div>
        </el-dialog>

        <el-dialog :visible.sync="dialogCenterForbidIpVisible" title="中心封ip">
          <el-form ref="dataForm" :model="monitorTemp" label-position="left" label-width="120px" style="width: 400px; margin-left:50px;">
            <el-form-item label="封号用户名">
              <el-input v-model="monitorTemp.playerName" :disabled="true"/>
            </el-form-item>
            <el-form-item label="封号IP">
              <el-input v-model="monitorTemp.ip" :disabled="true"/>
            </el-form-item>
            <el-form-item label="原因">
              <el-input v-model="monitorTemp.reason"/>
            </el-form-item>
            <el-form-item label="时长">
              <el-select v-model="monitorTemp.forbidTime" placeholder="时长" style="width: 120px" class="filter-item">
                <el-option v-for="item in chatForbidTimeArray" :key="item.key" :label="item.name" :value="item.key"/>
              </el-select>
            </el-form-item>
          </el-form>
          <div slot="footer" class="dialog-footer">
            <el-button @click="dialogCenterForbidIpVisible = false">取消</el-button>
            <el-button type="primary" @click="updateCenterIpForbid">确定</el-button>
          </div>
        </el-dialog>

        <el-dialog :visible.sync="dialogOperateFormVisible" title="聊天操作">
          <el-form ref="dataForm" :model="monitorChatTemp" label-position="left" label-width="120px" style="width: 400px; margin-left:50px;">
            <el-form-item label="用户名">
              <el-input v-model="monitorChatTemp.playerName" :disabled="true"/>
            </el-form-item>
            <el-form-item label="聊天内容">
              <span v-html="monitorChatTemp.chatMsg"/>
            </el-form-item>
            <el-form-item label="原因">
              <el-input v-model="monitorChatTemp.reason"/>
            </el-form-item>
            <el-form-item label="时长">
              <el-select v-model="monitorChatTemp.forbidTime" placeholder="时长" style="width: 120px" class="filter-item">
                <el-option v-for="item in chatForbidTimeArray" :key="item.key" :label="item.name" :value="item.key"/>
              </el-select>
            </el-form-item>
          </el-form>
          <div slot="footer" class="dialog-footer">
            <el-button @click="dialogOperateFormVisible = false">取消</el-button>
            <el-button type="primary" @click="updateForbidPlayerMonitor">封禁</el-button>
            <el-button type="primary" @click="updateForbidChatPlayerMonitor">禁言</el-button>
            <el-button type="primary" @click="updateIgnoreChatPlayerMonitor">禁默</el-button>
            <el-button type="primary" @click="updateKickOutPlayerMonitor">踢人</el-button>
            <el-button type="primary" @click="updateCenterUserForbidMonitor">中心封号</el-button>
            <el-button type="primary" @click="updateCenterIpForbidMonitor">中心封ip</el-button>
          </div>
        </el-dialog>

      </el-tab-pane>
      <el-tab-pane v-permission="['super_admin', 'super_channel', 'channel', 'platform', 'service', 'minitor','super_channel_service','common_service','gaoji_service']" label="封禁列表">
        <FengJinPanel></FengJinPanel>
      </el-tab-pane>
      <!-- <el-tab-pane label="封IP列表">封IP列表</el-tab-pane> -->
      <el-tab-pane v-permission="['super_admin', 'super_channel', 'channel', 'platform', 'service', 'minitor','super_channel_service','common_service','gaoji_service']" label="禁言列表">
        <JinYanPanel></JinYanPanel>
      </el-tab-pane>

     <el-tab-pane  v-permission="['super_admin', 'super_channel', 'channel', 'platform', 'service', 'minitor','super_channel_service','common_service','gaoji_service']" label="禁默列表">
        <JinMoPanel></JinMoPanel>
      </el-tab-pane>
      <el-tab-pane label="聊天日志">
        <PlayerChatLog></PlayerChatLog>
      </el-tab-pane>
    </el-tabs>
  </div>
</template>
<script>
import waves from "@/directive/waves"; // 水波纹指令
import permission from "@/directive/permission/index.js"; // 权限判断指令
import { getAllPlatformList } from "@/api/platform";
import { getCenterServerList } from "@/api/center";
import { getSensitive, saveUserInfo } from "@/api/sensitive";
import { newWebSocket } from "@/websocket/websocket";
import { QueueDelegate } from "@/websocket/queuedelegate";
import { parseTime } from "@/utils/index";
import PlayerChatLog from "./chatlog.vue";
import JinMoPanel from "./chat/jinmo.vue";
import FengJinPanel from "./chat/fengjin.vue";
import JinYanPanel from "./chat/jinyan.vue";
import {
  getFengJinPlayerList,
  getJinYanPlayerList,
  forbidPlayer,
  unForbidPlayer,
  forbidChatPlayer,
  unForbidChatPlayer,
  ignoreChatPlayer,
  unIgnoreChatPlayer,
  getJinMoPlayerList,
  getPlayerInfo,
  kickOutPlayer
} from "@/api/player";
import {
  getCenterUserInfo,
  updateCenterUserName,
  updateCenterForbid,
  updateCenterIpForbid
} from "@/api/centeruser";
import { chatTypeList, chatMethodList, chatForbidTimeList } from "@/types/chat";
import { makeSensitiveMap, replaceSensitiveWord } from "@/utils/sensitive";
import basemessage from "@/proto/basic_pb";
import chat_pb from "@/proto/chat_pb";
import chatmessagetype_pb from "@/proto/chatmessagetype_pb";
import { Message, MessageBox } from "element-ui";

export default {
  name: "MonitorList",
  components: {
    PlayerChatLog,
    JinMoPanel,
    FengJinPanel,
    JinYanPanel
  },
  directives: {
    waves,
    permission
  },
  filters: {
    parseTime: function(value) {
      return parseTime(value, "{y}-{m}-{d} {h}:{i}:{s}");
    },
    parseTimeSp: function(value) {
      if (!value) {
        return "永久";
      }
      return parseTime(value, "{y}-{m}-{d} {h}:{i}:{s}");
    },
    parseJin: function(value) {
      if (value == 1) {
        return "已封禁";
      }
      if (value == 0) {
        return "已解封";
      }
    },
    chatTypeFilter: function(value) {
      return chatTypeList[value].name;
    },
    chatMethodFilter: function(value) {
      return chatMethodList[value].name;
    }
  },
  data() {
    return {
      listLoading: false,
      chatKey: 0,
      total: 0,
      //元数据
      chatForbidTimeArray: [],
      chatListQuery: {
        pageIndex: 1,
        platformId: undefined,
        serverArray: [],
        oldServerArray: [],
        channelName: "",
        playerName: undefined,
        toPlayerName: undefined,
        vipLevel: undefined,
        chatMsg: undefined,
        socketbtnType: "primary",
        socketbtnMsg: "启动",
        socketbtnState: false
      },
      sensitive: {
        content: ""
      },
      dialogFormVisible: false,
      dialogForbidFormVisible: false,
      dialogForbidChatFormVisible: false,
      dialogIgnoreChatFormVisible: false,
      dialogKickVisible:false,
      dialogCenterForbidVisible:false,
      dialogCenterForbidIpVisible:false,

      dialogOperateFormVisible: false,
      websocket: undefined,
      platformList: [],
      serverList: [],
      minGanCiList: [],
      sensitiveMap: undefined,
      chatTypeList: [],
      chatList: [],
      chatState: false,

      // 聊天监控里的行
      monitorTemp: {},
      //聊天监控里异常的行
      monitorChatTemp: {},

      // 封禁tab开始
      fengjinTableKey: 1,
      fengJinListQuery: {
        platformId: undefined,
        serverId: undefined,
        playerName: undefined,
        reason: undefined,
        pageIndex: 1,
        centerPlatformId: undefined,
        centerServerId: undefined
      },
      fengjinServerList: [],
      fengjinUserList: [],
      fengjinUserCount: 0,
      unForbidTemp: {}, // 解禁用户传入对象
      dialogUnForbidFormVisible: false,

      // 禁言tab开始
      jinyanTableKey: 1,
      jinyanListQuery: {
        platformId: undefined,
        serverId: undefined,
        playerName: undefined,
        reason: undefined,
        pageIndex: 1,
        centerPlatformId: undefined,
        centerServerId: undefined
      },
      jinyanServerList: [],
      jinyanUserList: [],
      jinyanUserCount: 0,
      unForbidChatTemp: {},
      dialogUnForbidChatFormVisible: false,

      // 禁莫tab开始
      jinMoTableKey: 1,
      jinMoListQuery: {
        platformId: undefined,
        serverId: undefined,
        playerName: undefined,
        reason: undefined,
        pageIndex: 1,
        centerPlatformId: undefined,
        centerServerId: undefined
      },
      jinMoServerList: [],
      jinMoUserList: [],
      jinMoUserCount: 0,
      unJinMoChatTemp: {},
      dialogUnJinMoChatFormVisible: false,

      autoIndex: 0,
      checkSocketState: true
    };
  },
  created() {
    this.initMetaData();
    // this.handleScoket();
  },
  mounted() {
    //   this.websocket.close()
  },
  beforeDestroy() {
    if (this.websocket ){
      this.checkSocketState = false; // 停止状态检测
      this.websocket.close();
    }
  },
  destroyed() {},
  methods: {
    handleFilter: function() {
      if (!this.chatListQuery.socketbtnState) {
        this.chatListQuery.socketbtnState = true;
        this.chatListQuery.socketbtnMsg = "停止";
        this.chatListQuery.socketbtnType = "danger";
          this.checkSocketState = true;
        if (!this.websocket || this.websocket.readyState != 1) {
          this.handleScoket();
          setTimeout(() => {
            this.sendPlatChange();
          }, 1000);
          return;
        }
        this.sendPlatChange();
        // this.autoAddChat();
        return;
      }
      this.checkSocketState = false;
       this.websocket.close();
      // this.sendEmptyPlatChange();
      this.chatListQuery.socketbtnState = false;
      this.chatListQuery.socketbtnMsg = "启动";
      this.chatListQuery.socketbtnType = "primary";
    },
    // websocket断线重连,在错误或者关闭的时候
    startWebScoket() {
      if (!this.checkSocketState) {
        return;
      }
      if (!this.websocket ||  this.websocket.readyState != 1 && this.websocket.readyState != 0) {
        this.handleScoket();
        setTimeout(() => {
          this.sendPlatChange();
        }, 1000);
      }
    },
    handleSearch(){},
    sendPlatChange() {
      // console.log(this.chatListQuery.serverArray);
      if (this.websocket.readyState != 1) {
        return;
      }
        console.log("monitor")
       console.log(this.chatListQuery.serverArray)
      
      let tempServerList = [];
      for (const item of this.chatListQuery.serverArray) {
        if (item === 0) {
          continue;
        }
        tempServerList.push(item);
      }
      const chatmsg = new chat_pb.CGChatMinitor();
      chatmsg.setServerlistList(tempServerList);
      const msg = new basemessage.Message();
      msg.setMessagetype(
        chatmessagetype_pb.ChatMonitorMessageType.CGCHATMINITORTYPE
      );
      msg.setExtension(chat_pb.cgchatminitor, chatmsg);
      // console.log("发送玩家服务配置");
      // console.log(msg);
      this.websocket.send(msg.serializeBinary());
    },
    sendEmptyPlatChange() {
      if (this.websocket.readyState != 1) {
        return;
      }
      const chatmsg = new chat_pb.CGChatMinitor();
      chatmsg.setServerlistList([]);
      const msg = new basemessage.Message();
      msg.setMessagetype(
        chatmessagetype_pb.ChatMonitorMessageType.CGCHATMINITORTYPE
      );
      msg.setExtension(chat_pb.cgchatminitor, chatmsg);
      // console.log("发送玩家空服务配置");
      //   console.log(msg);
      this.websocket.send(msg.serializeBinary());
    },
    initMetaData() {
      this.chatTypeList = chatTypeList;
      this.chatForbidTimeArray = chatForbidTimeList;
      getAllPlatformList().then(res => {
        this.platformList = res.itemArray;
      });
      getSensitive().then(res => {
        this.sensitive.content = res.content;
        this.spilitSensitive();
      });
    },
    handlePlatformChange(e) {
      const item = this.findPlatFormItem(e);
      if (item) {
        getCenterServerList(item.centerPlatformId).then(res => {
          this.serverList = res.itemArray;
          if (this.serverList.length > 0) {
            let allSelect = { id: 0, serverName: "全选" };
            this.serverList.unshift(allSelect);
            this.chatListQuery.oldServerArray = [];
          }
        });
      }
    },
    handleServerChange(e) {
      this.selectAll(e);
      console.log(this.serverList);
      if (this.chatListQuery.socketbtnState) {
        this.sendPlatChange();
      }
    },
    selectAll(val) {
      const allValues = [];
      // 保留所有值
      for (const item of this.serverList) {
        allValues.push(item.id);
      }
      // 用来储存上一次的值，可以进行对比
      const oldVal =
        this.chatListQuery.oldServerArray.length === 1
          ? this.chatListQuery.oldServerArray[0]
          : [];
      // 若是全部选择
      if (val.includes(0)) this.chatListQuery.serverArray = allValues;
      // 取消全部选中 上次有 当前没有 表示取消全选
      if (oldVal.includes(0) && !val.includes(0))
        this.chatListQuery.serverArray = [];
      // 点击非全部选中 需要排除全部选中 以及 当前点击的选项
      // 新老数据都有全部选中
      if (oldVal.includes(0) && val.includes(0)) {
        const index = val.indexOf(0);
        val.splice(index, 1); // 排除全选选项
        this.chatListQuery.serverArray = val;
      }
      // 全选未选 但是其他选项全部选上 则全选选上 上次和当前 都没有全选
      if (!oldVal.includes(0) && !val.includes(0)) {
        if (val.length === allValues.length - 1)
          this.chatListQuery.serverArray = [0].concat(val);
      }
      // 储存当前最后的结果 作为下次的老数据
      this.chatListQuery.oldServerArray[0] = this.chatListQuery.serverArray;
    },
    findPlatFormItem(platformId) {
      const platform = this.platformList.find(n => {
        return n.platformId == platformId;
      });
      if (platform) {
        return platform;
      }
      return undefined;
    },
    handleScoket: function() {
      this.queuedelegate = new QueueDelegate();
      this.queuedelegate.register(
        chatmessagetype_pb.ChatMonitorMessageType.GCCHATMINITORMSGTYPE,
        this.handleSocketChatMsg
      );
      this.queuedelegate.register(
        chatmessagetype_pb.ChatMonitorMessageType.GCCHATMINITORTYPE,
        this.handleScoketSetUserServer
      );
      const socketEvent = {};
      socketEvent.onclose = this.startWebScoket;
      const websocketAPI = process.env.WEBSOCKET_API + "/websocket";
      console.log(websocketAPI);
      this.websocket = newWebSocket(
        websocketAPI, // 这个之后要写配置
        this.queuedelegate,
        socketEvent
      );
    },
    handleMinGan: function() {
      this.dialogFormVisible = true;
      // this.websocket.close();
    },
    autoAddChat: function() {
      setTimeout(() => {
        this.autoIndex = this.autoIndex + 1;
        const item = {
          playerName: "playerName" + this.autoIndex,
          vipLevel: this.autoIndex,
          playerLevel: 2,
          chatType: "世界",
          toPlayerName: "我的小朋友" + this.autoIndex,
          chatMethod: "文字",
          chatMsg:
            '小朋友上课要认真听讲，是，是，是，，是，是，是，，<span style="background:#E6A23C">算是</span>，' +
            this.autoIndex,
          chatTime: "2019-05-05"
        };
        // let secondList = [item]
        // secondList.concat(this.chatList)
        this.chatList.unshift(item);
        // this.chatList = secondList
        if (this.chatList.length > 100) {
          this.chatList = this.chatList.slice(0, 100);
        }
        if (this.chatState) {
          this.autoAddChat();
        }
      }, 20);
    },
    handleSocketChatMsg: function(e) {
      if (!e) {
        return;
      }
      const data = e.getExtension(chat_pb.gcchatminitormsg);
      if (!data) {
        return;
      }

      const playerid = data.getPlayerid();
      const playerName = data.getPlayername();
      const vipLevel = data.getViplevel();
      const gameLevel = data.getGamelevel();
      const chatType = data.getChattype();
      const chatMethod = data.getChatmethod();
      let chatMsg = data.getChatmsg();
      const chatTime = data.getChattime();
      const toPlayerName = data.getToplayername();
      const ip = data.getIp();
      const centerPlatformId = data.getCenterplatformid();
      const centerServerId = data.getCenterserverid();
      const userId = data.getUserid();

      // 会话数据的筛选
      // 玩家名字筛选
      if (this.chatListQuery.playerName) {
        if (playerName.indexOf(this.chatListQuery.playerName) < 0) {
          return;
        }
      }

      if (this.chatListQuery.toPlayerName) {
        if (toPlayerName.indexOf(this.chatListQuery.toPlayerName) < 0) {
          return;
        }
      }

      if (this.chatListQuery.vipLevel) {
        if (parseInt(vipLevel) != parseInt(this.chatListQuery.vipLevel)) {
          // console.log(parseInt(vipLevel));
          // console.log(parseInt(chatListQuery.vipLevel));
          return;
        }
      }

      if (this.chatListQuery.chatType) {
        if (parseInt(chatType) != parseInt(this.chatListQuery.chatType)) {
          return;
        }
      }

      if (this.chatListQuery.chatMsg) {
        if (chatMsg.indexOf(this.chatListQuery.chatMsg) < 0) {
          return;
        }
      }
      let preChatMsg = chatMsg;
      if (chatMsg && this.minGanCiList && this.minGanCiList.length > 0) {
        chatMsg = replaceSensitiveWord(
          this.sensitiveMap,
          chatMsg,
          '<span style="background:#E6A23C">',
          "</span>"
        );

        // console.log('chatmsg')
        // console.log(chatMsg)

        // for (let i = 0, len = this.minGanCiList.length; i < len; i++) {
        //   let mingan = this.minGanCiList[i];
        //   let newMingan = replaceSensitiveWord(this.sen);
        //   '<span style="background:#E6A23C">' + mingan + "</span>";
        //   chatMsg = chatMsg.replace(mingan, newMingan);
        // }
      }

      const item = {
        playerName: playerName,
        vipLevel: vipLevel,
        playerLevel: gameLevel,
        chatType: chatType,
        toPlayerName: toPlayerName,
        chatMethod: chatMethod,
        chatMsg: chatMsg,
        chatTime: chatTime,
        ip: ip,
        centerServerId: centerServerId,
        centerPlatformId: centerPlatformId,
        playerId: playerid,
        userId:userId
      };
      this.chatList.unshift(item);
      if (this.chatList.length > 100) {
        this.chatList = this.chatList.slice(0, 100);
      }
      if (preChatMsg.length != chatMsg.length) {
        this.dialogOperateFormVisible = true;
        this.monitorChatTemp = item;
      }
    },
    updateSensitive() {
      saveUserInfo(this.sensitive).then(res => {
        this.dialogFormVisible = false;
        this.spilitSensitive();
        this.showSuccess();
      });
    },
    handleScoketSetUserServer: function(e) {
      console.log("处理队列设置回复");
    },
    showSuccess() {
      this.$message({
        message: "设置成功",
        type: "success",
        duration: 1000
      });
    },
    spilitSensitive() {
      this.minGanCiList = this.sensitive.content.split(",");
      this.sensitiveMap = makeSensitiveMap(this.minGanCiList);
    },

    // 封禁用户
    handleForbid(e) {
      this.monitorTemp = e;
      this.monitorTemp.reason = undefined;
      this.dialogForbidFormVisible = true;
    },
    updateForbidPlayer(e) {
      const postData = {
        centerPlatformId: this.monitorTemp.centerPlatformId,
        centerServerId: this.monitorTemp.centerServerId,
        playerId: this.monitorTemp.playerId,
        reason: this.monitorTemp.reason,
        forbidTime: this.monitorTemp.forbidTime
      };

      console.log(postData);
      forbidPlayer(postData).then(res => {
        this.dialogForbidFormVisible = false;
        this.dialogOperateFormVisible = false;
        this.showSuccess();
      });
    },
    updateForbidPlayerMonitor(e) {
      const postData = {
        centerPlatformId: this.monitorChatTemp.centerPlatformId,
        centerServerId: this.monitorChatTemp.centerServerId,
        playerId: this.monitorChatTemp.playerId,
        reason: this.monitorChatTemp.reason,
        forbidTime: this.monitorChatTemp.forbidTime
      };

      console.log(postData);
      forbidPlayer(postData).then(res => {
        this.dialogOperateFormVisible = false;
        this.showSuccess();
      });
    },

    // 禁言用户
    handleForbidChat(e) {
      this.monitorTemp = e;
      this.monitorTemp.reason = undefined;
      this.dialogForbidChatFormVisible = true;
    },
    updateForbidChatPlayer(e) {
      const postData = {
        centerPlatformId: this.monitorTemp.centerPlatformId,
        centerServerId: this.monitorTemp.centerServerId,
        playerId: this.monitorTemp.playerId,
        reason: this.monitorTemp.reason,
        forbidTime: this.monitorTemp.forbidTime
      };

      forbidChatPlayer(postData).then(res => {
        this.dialogForbidChatFormVisible = false;
        this.dialogOperateFormVisible = false;
        this.showSuccess();
      });
    },
    updateForbidChatPlayerMonitor(e) {
      const postData = {
        centerPlatformId: this.monitorChatTemp.centerPlatformId,
        centerServerId: this.monitorChatTemp.centerServerId,
        playerId: this.monitorChatTemp.playerId,
        reason: this.monitorChatTemp.reason,
        forbidTime: this.monitorChatTemp.forbidTime
      };

      forbidChatPlayer(postData).then(res => {
        this.dialogOperateFormVisible = false;
        this.showSuccess();
      });
    },

    //禁默用户
    handleIgnoreChat(e) {
      this.monitorTemp = e;
      this.monitorTemp.reason = undefined;
      this.dialogIgnoreChatFormVisible = true;
    },
    updateIgnoreChatPlayer(e) {
      const postData = {
        centerPlatformId: this.monitorTemp.centerPlatformId,
        centerServerId: this.monitorTemp.centerServerId,
        playerId: this.monitorTemp.playerId,
        reason: this.monitorTemp.reason,
        forbidTime: this.monitorTemp.forbidTime
      };

      ignoreChatPlayer(postData).then(res => {
        this.dialogIgnoreChatFormVisible = false;
        this.dialogOperateFormVisible = false;
        this.showSuccess();
      });
    },
    //踢人
    handleKickOut(e){
      this.monitorTemp = e;
      this.monitorTemp.reason = undefined;
      this.dialogKickVisible = true;
    },
    updateKickOutPlayer(e) {
      const postData = {
        centerPlatformId: this.monitorTemp.centerPlatformId,
        centerServerId: this.monitorTemp.centerServerId,
        playerId: this.monitorTemp.playerId,
        reason: this.monitorTemp.reason
      };
      console.log(postData);
      kickOutPlayer(postData).then(res => {
        this.dialogKickVisible = false;
        this.showSuccess();
      });
    },
    //封号
    handleCenterForbid(e) {
      this.monitorTemp = e;
      this.monitorTemp.reason = undefined;
      this.dialogCenterForbidVisible = true;
    },
    updateCenterUserForbid(e) {
      const postData = {
        userId: parseInt(this.monitorTemp.userId),
        reason: this.monitorTemp.reason,
        forbid: 1,
        forbidTime: this.monitorTemp.forbidTime
      };
      console.log(postData);
      updateCenterForbid(postData).then(res => {
        this.dialogCenterForbidVisible = false;
        this.showSuccess();
      });
    },
    //封IP
    handleIpForbid(e) {
      this.monitorTemp = e;
      this.monitorTemp.reason = undefined;
      this.dialogCenterForbidIpVisible = true;
    },
    updateCenterIpForbid(e) {
      const postData = {
        ip: this.monitorTemp.ip,
        reason: this.monitorTemp.reason,
        forbid: 1,
        forbidTime: this.monitorTemp.forbidTime
      };
      console.log(postData);
      updateCenterIpForbid(postData).then(res => {
        this.dialogCenterForbidIpVisible = false;
        this.showSuccess();
      });
    },
    updateIgnoreChatPlayerMonitor(e) {
      const postData = {
        centerPlatformId: this.monitorChatTemp.centerPlatformId,
        centerServerId: this.monitorChatTemp.centerServerId,
        playerId: this.monitorChatTemp.playerId,
        reason: this.monitorChatTemp.reason,
        forbidTime: this.monitorChatTemp.forbidTime
      };

      ignoreChatPlayer(postData).then(res => {
        this.dialogOperateFormVisible = false;
        this.showSuccess();
      });
    },
    updateKickOutPlayerMonitor(e) {
      const postData = {
        centerPlatformId: this.monitorChatTemp.centerPlatformId,
        centerServerId: this.monitorChatTemp.centerServerId,
        playerId: this.monitorChatTemp.playerId,
        reason: this.monitorChatTemp.reason
      };
      console.log(postData);
      kickOutPlayer(postData).then(res => {
        this.dialogOperateFormVisible = false;
        this.showSuccess();
      });
    },
    updateCenterUserForbidMonitor(e) {
      const postData = {
        userId: parseInt(this.monitorChatTemp.userId),
        reason: this.monitorChatTemp.reason,
        forbid: 1,
        forbidTime: this.monitorChatTemp.forbidTime
      };
      console.log(postData);
      updateCenterForbid(postData).then(res => {
        this.dialogOperateFormVisible = false;
        this.showSuccess();
      });
    },
    updateCenterIpForbidMonitor(e) {
      const postData = {
        ip: this.monitorChatTemp.ip,
        reason: this.monitorChatTemp.reason,
        forbid: 1,
        forbidTime: this.monitorChatTemp.forbidTime,
        centerPlatformId: this.monitorChatTemp.centerPlatformId,
        centerServerId: this.monitorChatTemp.centerServerId,
        playerId: this.monitorChatTemp.playerId
      };
      console.log(postData);
      updateCenterIpForbid(postData).then(res => {
        this.dialogOperateFormVisible = false;
        this.showSuccess();
      });
    },
  }
};
</script>
