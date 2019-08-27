<template>
  <div class="app-container">
    <el-tabs type="border-card">
      <el-tab-pane label="聊天监控">
        <div>
          <div class="filter-container">
            <el-select v-model="chatListQuery.platformId" placeholder="平台" style="width: 160px" class="filter-item" @change="handlePlatformChange">
              <el-option v-for="item in platformList" :key="item.platformId" :label="item.platformName" :value="item.platformId" />
            </el-select>
            <el-select v-model="chatListQuery.serverArray" multiple collapse-tags placeholder="服务器" clearable style="width: 180px" class="filter-item" @change="handleServerChange">
              <el-option v-for="item in serverList" :key="item.id" :label="item.serverName" :value="item.id"/>
            </el-select>
            <el-input v-model.trim="chatListQuery.playerName" placeholder="角色名" style="width: 200px;" class="filter-item" />
            <el-input v-model.trim="chatListQuery.toPlayerName" placeholder="私聊目标" style="width: 200px;" class="filter-item" />
            <el-input v-model="chatListQuery.vipLevel" placeholder="VIP等级" style="width: 200px;" class="filter-item" />

            <el-select v-model="chatListQuery.chatType" placeholder="类型" clearable style="width: 120px" class="filter-item">
              <el-option v-for="item in chatTypeList" :key="item.key" :label="item.name" :value="item.key"/>
            </el-select>
            <el-input v-model.trim="chatListQuery.chatMsg" placeholder="聊天内容" style="width: 200px;" class="filter-item" />

            <el-button v-waves :type="chatListQuery.socketbtnType" class="filter-item" icon="el-icon-search" @click="handleFilter">{{ chatListQuery.socketbtnMsg }}</el-button>
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
          <el-table-column label="角色名" align="center" width="180px">
            <template slot-scope="scope">
              <span>{{ scope.row.playerName }}</span>
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
          <el-table-column label="类型" width="60px">
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
          <el-table-column label="操作" align="center" width="260" class-name="small-padding fixed-width">
            <template slot-scope="scope">
              <el-button type="primary" size="mini" @click="handleForbid(scope.row)">封禁</el-button>
              <!-- <el-button type="primary" size="mini" @click="handlePwd(scope.row)">封IP</el-button> -->
              <el-button size="mini" type="danger" @click="handleForbidChat(scope.row)">禁言</el-button>
              <el-button size="mini" type="danger" @click="handleIgnoreChat(scope.row)">禁默</el-button>
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

      </el-tab-pane>
      <el-tab-pane label="封禁列表">
        

      </el-tab-pane>
      <!-- <el-tab-pane label="封IP列表">封IP列表</el-tab-pane> -->
      <el-tab-pane label="禁言列表">
        <div>
          <div class="filter-container">
            <el-select v-model="jinyanListQuery.platformId" placeholder="平台" style="width: 160px" class="filter-item" @change="handleJinYanPlatformChange">
              <el-option v-for="item in platformList" :key="item.platformId" :label="item.platformName" :value="item.platformId" />
            </el-select>
            <el-select v-model="jinyanListQuery.serverId" collapse-tags placeholder="服务器" clearable style="width: 180px" class="filter-item" @change="handleJinYanServerChange">
              <el-option v-for="item in jinyanServerList" :key="item.id" :label="item.serverName" :value="item.id"/>
            </el-select>
            <el-input v-model.trim="jinyanListQuery.playerName" placeholder="角色名" style="width: 200px;" class="filter-item" />
            <el-input v-model.trim="jinyanListQuery.reason" placeholder="禁言理由" style="width: 200px;" class="filter-item" />
            <el-button v-waves class="filter-item" type="primary" icon="el-icon-search" @click="handleJinYanFilter">搜索</el-button>
          </div>
        </div>
        <el-table
          v-loading="listLoading"
          :key="jinyanTableKey"
          :data="jinyanUserList"
          border
          fit
          highlight-current-row
          style="width: 100%;margin-top:15px;">
          <el-table-column label="角色ID" align="center" width="180px">
            <template slot-scope="scope">
              <span>{{ scope.row.id }}</span>
            </template>
          </el-table-column>
          <el-table-column label="账户ID" align="center" width="180px">
            <template slot-scope="scope">
              <span>{{ scope.row.playerId }}</span>
            </template>
          </el-table-column>
          <el-table-column label="角色名" align="center" width="180px">
            <template slot-scope="scope">
              <span>{{ scope.row.playerName }}</span>
            </template>
          </el-table-column>
          <el-table-column label="禁言状态" width="250px" align="center">
            <template slot-scope="scope">
              <span v-if="scope.row.forbidChat == 1" style="color:#F56C6C">{{ scope.row.forbidChat | parseJin }}</span>
              <span v-else style="color:#67C23A">{{ scope.row.forbidChat | parseJin }}</span>
            </template>
          </el-table-column>
          <el-table-column label="禁言理由" width="200px">
            <template slot-scope="scope">
                <span style="color:#F56C6C;font-weight:bold;}">{{ scope.row.forbidChatText }}</span>
            </template>
          </el-table-column>
          <el-table-column label="禁言时间" width="180px">
            <template slot-scope="scope">
              <span>{{ scope.row.forbidChatTime | parseTime }}</span>
            </template>
          </el-table-column>
          <el-table-column label="解禁时间" width="180px">
            <template slot-scope="scope">
              <span>{{ scope.row.forbidChatEndTime | parseTimeSp }}</span>
            </template>
          </el-table-column>
          <el-table-column label="禁言者" min-width="120px">
            <template slot-scope="scope">
              <span>{{ scope.row.forbidChatName }}</span>
            </template>
          </el-table-column>
          <el-table-column label="操作" align="center" width="260" class-name="small-padding fixed-width">
            <template slot-scope="scope">
              <el-button v-if="scope.row.forbidChat == 1" type="primary" size="mini" @click="handleJieJinYan(scope.row)">解封</el-button>
            </template>
          </el-table-column>
        </el-table>
        <div class="pagination-container" style="margin-top:15px;">
          <el-pagination :current-page="jinyanListQuery.pageIndex" :page-sizes="[20]" :total="jinyanUserCount" background layout="total, sizes, prev, pager, next, jumper" @current-change="handleJinYanCurrentChange"/>
        </div>

        <el-dialog :visible.sync="dialogUnForbidChatFormVisible" title="是否解禁用户">
          <el-form ref="dataForm" :model="unForbidChatTemp" label-position="left" label-width="120px" style="width: 400px; margin-left:50px;">
            <el-form-item label="封禁用户名">
              <el-input v-model="unForbidChatTemp.playerName" :disabled="true"/>
            </el-form-item>
          </el-form>
          <div slot="footer" class="dialog-footer">
            <el-button @click="dialogUnForbidChatFormVisible = false">取消</el-button>
            <el-button type="primary" @click="updateUnForbidChatPlayer">解禁</el-button>
          </div>
        </el-dialog>
      </el-tab-pane>

     <el-tab-pane label="禁默列表">
        <div>
          <div class="filter-container">
            <el-select v-model="jinMoListQuery.platformId" placeholder="平台" style="width: 160px" class="filter-item" @change="handleJinMoPlatformChange">
              <el-option v-for="item in platformList" :key="item.platformId" :label="item.platformName" :value="item.platformId" />
            </el-select>
            <el-select v-model="jinMoListQuery.serverId" collapse-tags placeholder="服务器" clearable style="width: 180px" class="filter-item" @change="handleJinMoServerChange">
              <el-option v-for="item in jinMoServerList" :key="item.id" :label="item.serverName" :value="item.id"/>
            </el-select>
            <el-input v-model.trim="jinMoListQuery.playerName" placeholder="角色名" style="width: 200px;" class="filter-item" />
            <el-input v-model.trim="jinMoListQuery.reason" placeholder="禁默理由" style="width: 200px;" class="filter-item" />
            <el-button v-waves class="filter-item" type="primary" icon="el-icon-search" @click="handleJinMoFilter">搜索</el-button>
          </div>
        </div>
        <el-table
          v-loading="listLoading"
          :key="jinMoTableKey"
          :data="jinMoUserList"
          border
          fit
          highlight-current-row
          style="width: 100%;margin-top:15px;">
          <el-table-column label="角色ID" align="center" width="180px">
            <template slot-scope="scope">
              <span>{{ scope.row.id }}</span>
            </template>
          </el-table-column>
          <el-table-column label="账户ID" align="center" width="180px">
            <template slot-scope="scope">
              <span>{{ scope.row.playerId }}</span>
            </template>
          </el-table-column>
          <el-table-column label="角色名" align="center" width="180px">
            <template slot-scope="scope">
              <span>{{ scope.row.playerName }}</span>
            </template>
          </el-table-column>
          <el-table-column label="禁默状态" width="250px" align="center">
            <template slot-scope="scope">
              <span v-if="scope.row.ignoreChat == 1" style="color:#F56C6C">{{ scope.row.ignoreChat | parseJin }}</span>
              <span v-else style="color:#67C23A">{{ scope.row.ignoreChat | parseJin }}</span>
            </template>
          </el-table-column>
          <el-table-column label="禁默理由" width="200px">
            <template slot-scope="scope">
                <span style="color:#F56C6C;font-weight:bold;}">{{ scope.row.ignoreChatText }}</span>
            </template>
          </el-table-column>
          <el-table-column label="禁默时间" width="180px">
            <template slot-scope="scope">
              <span>{{ scope.row.ignoreChatTime | parseTime }}</span>
            </template>
          </el-table-column>
          <el-table-column label="解禁时间" width="180px">
            <template slot-scope="scope">
              <span>{{ scope.row.ignoreChatEndTime | parseTimeSp }}</span>
            </template>
          </el-table-column>
          <el-table-column label="禁言者" min-width="120px">
            <template slot-scope="scope">
              <span>{{ scope.row.ignoreChatName }}</span>
            </template>
          </el-table-column>
          <el-table-column label="操作" align="center" width="260" class-name="small-padding fixed-width">
            <template slot-scope="scope">
              <el-button v-if="scope.row.ignoreChat == 1" type="primary" size="mini" @click="handleJieJinMo(scope.row)">解封</el-button>
            </template>
          </el-table-column>
        </el-table>
        <div class="pagination-container" style="margin-top:15px;">
          <el-pagination :current-page="jinMoListQuery.pageIndex" :page-sizes="[20]" :total="jinMoUserCount" background layout="total, sizes, prev, pager, next, jumper" @current-change="handleJinMoCurrentChange"/>
        </div>

        <el-dialog :visible.sync="dialogUnJinMoChatFormVisible" title="是否解禁用户">
          <el-form ref="dataForm" :model="unJinMoChatTemp" label-position="left" label-width="120px" style="width: 400px; margin-left:50px;">
            <el-form-item label="封禁用户名">
              <el-input v-model="unJinMoChatTemp.playerName" :disabled="true"/>
            </el-form-item>
          </el-form>
          <div slot="footer" class="dialog-footer">
            <el-button @click="dialogUnJinMoChatFormVisible = false">取消</el-button>
            <el-button type="primary" @click="updateUnJinMoChatPlayer">解禁</el-button>
          </div>
        </el-dialog>
      </el-tab-pane>
      <el-tab-pane label="聊天日志">
        <PlayerChatLog></PlayerChatLog>
      </el-tab-pane>
    </el-tabs>
  </div>
</template>
<script>
import waves from "@/directive/waves"; // 水波纹指令
import { getAllPlatformList } from "@/api/platform";
import { getCenterServerList } from "@/api/center";
import { getSensitive, saveUserInfo } from "@/api/sensitive";
import { newWebSocket } from "@/websocket/websocket";
import { QueueDelegate } from "@/websocket/queuedelegate";
import { parseTime } from "@/utils/index";
import PlayerChatLog  from "./chatlog.vue";
import {
  getFengJinPlayerList,
  getJinYanPlayerList,
  forbidPlayer,
  unForbidPlayer,
  forbidChatPlayer,
  unForbidChatPlayer,
  ignoreChatPlayer,
  unIgnoreChatPlayer,
  getJinMoPlayerList
} from "@/api/player";
import { chatTypeList, chatMethodList, chatForbidTimeList } from "@/types/chat";
import { makeSensitiveMap, replaceSensitiveWord } from "@/utils/sensitive";
import basemessage from "@/proto/basic_pb";
import chat_pb from "@/proto/chat_pb";
import chatmessagetype_pb from "@/proto/chatmessagetype_pb";
import { Message, MessageBox } from "element-ui";

export default {
  name: "MonitorList",
  components: {
    PlayerChatLog
  },
  directives: {
    waves
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
    this.handleScoket();
  },
  mounted() {
    //   this.websocket.close()
  },
  beforeDestroy() {
    this.checkSocketState = false; // 停止状态检测
    this.websocket.close();
  },
  destroyed() {},
  methods: {
    handleFilter: function() {
      if (!this.chatListQuery.socketbtnState) {
        this.chatListQuery.socketbtnState = true;
        this.chatListQuery.socketbtnMsg = "停止";
        this.chatListQuery.socketbtnType = "danger";
        if (this.websocket.readyState != 1) {
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
      this.sendEmptyPlatChange();
      this.chatListQuery.socketbtnState = false;
      this.chatListQuery.socketbtnMsg = "启动";
      this.chatListQuery.socketbtnType = "primary";
    },
    // websocket断线重连,在错误或者关闭的时候
    startWebScoket() {
      if (!this.checkSocketState) {
        return;
      }
      if (this.websocket.readyState != 1 && this.websocket.readyState != 0) {
        this.handleScoket();
        setTimeout(() => {
          this.sendPlatChange();
        }, 1000);
      }
    },
    sendPlatChange() {
      // console.log(this.chatListQuery.serverArray);
      if (this.websocket.readyState != 1) {
        return;
      }
      const chatmsg = new chat_pb.CGChatMinitor();
      chatmsg.setServerlistList(this.chatListQuery.serverArray);
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
        });
      }
    },
    handleServerChange(e) {
      if (this.chatListQuery.socketbtnState) {
        this.sendPlatChange();
      }
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
        playerId: playerid
      };
      this.chatList.unshift(item);
      if (this.chatList.length > 100) {
        this.chatList = this.chatList.slice(0, 100);
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
        this.showSuccess();
      });
    },

    /** **********************************封禁设置 ****************************/

    // 封禁设置
    handleFengjinPlatformChange(e) {
      const item = this.findPlatFormItem(e);
      if (item) {
        getCenterServerList(item.centerPlatformId).then(res => {
          this.fengjinServerList = res.itemArray;
        });
      }
    },
    handleFengJinServerChange(e) {
      const serverInfo = this.findFengJinServerItem(e);
      if (serverInfo) {
        this.fengJinListQuery.centerPlatformId = serverInfo.centerPlatformId;
        this.fengJinListQuery.centerServerId = serverInfo.serverId;
      }
    },
    findFengJinServerItem(serverId) {
      const server = this.fengjinServerList.find(n => {
        return n.id == serverId;
      });
      if (server) {
        return server;
      }
      return undefined;
    },

    handleFengJinFilter(e) {
      // 搜索
      if (
        !this.fengJinListQuery.centerPlatformId ||
        !this.fengJinListQuery.centerServerId
      ) {
        Message({
          message: "请选择服务器",
          type: "error",
          duration: 1.5 * 1000
        });
        return;
      }
      this.fengJinListQuery.pageIndex = 1;
      console.log(this.fengJinListQuery);

      this.loadFengJin();
    },
    handleFengJinCurrentChange(e) {
      // 分页
      this.fengJinListQuery.pageIndex = e;
      this.loadFengJin();
    },
    handleJieFengJin(e) {
      this.dialogUnForbidFormVisible = true;
      this.unForbidTemp = e;
      console.log("解除封禁");
      console.log(e);
      // 解除封禁
    },
    updateUnForbidPlayer(e) {
      const postdata = {
        centerPlatformId: this.unForbidTemp.centerPlatformId,
        centerServerId: this.unForbidTemp.centerServerId,
        playerId: this.unForbidTemp.id
      };
      unForbidPlayer(postdata).then(res => {
        this.dialogUnForbidFormVisible = false;
        this.showSuccess();
        setTimeout(() => {
          this.loadFengJin();
        }, 300);
        // this.loadFengJin();
      });
    },
    loadFengJin() {
      this.listLoading = true;
      getFengJinPlayerList(this.fengJinListQuery).then(res => {
        this.fengjinUserList = res.itemArray;
        this.fengjinUserCount = res.total;
        this.listLoading = false;
      });
    },
    /** ********************禁言设置********************* */
    // 禁言设置
    handleJinYanPlatformChange(e) {
      const item = this.findPlatFormItem(e);
      if (item) {
        getCenterServerList(item.centerPlatformId).then(res => {
          this.jinyanServerList = res.itemArray;
        });
      }
    },
    handleJinYanServerChange(e) {
      const serverInfo = this.findJinYanServerItem(e);
      if (serverInfo) {
        this.jinyanListQuery.centerPlatformId = serverInfo.centerPlatformId;
        this.jinyanListQuery.centerServerId = serverInfo.serverId;
      }
    },
    findJinYanServerItem(serverId) {
      const server = this.jinyanServerList.find(n => {
        return n.id == serverId;
      });
      if (server) {
        return server;
      }
      return undefined;
    },

    handleJinYanFilter(e) {
      // 搜索
      if (
        !this.jinyanListQuery.centerPlatformId ||
        !this.jinyanListQuery.centerServerId
      ) {
        Message({
          message: "请选择服务器",
          type: "error",
          duration: 1.5 * 1000
        });
        return;
      }
      this.jinyanListQuery.pageIndex = 1;

      this.loadJinYan();
    },
    handleJinYanCurrentChange(e) {
      // 分页
      this.jinyanListQuery.pageIndex = e;
      this.loadJinYan();
    },
    handleJieJinYan(e) {
      // 解除禁言
      this.dialogUnForbidChatFormVisible = true;
      this.unForbidChatTemp = e;
      console.log(e);
    },
    updateUnForbidChatPlayer(e) {
      const postdata = {
        centerPlatformId: this.unForbidChatTemp.centerPlatformId,
        centerServerId: this.unForbidChatTemp.centerServerId,
        playerId: this.unForbidChatTemp.id
      };
      unForbidChatPlayer(postdata).then(res => {
        this.dialogUnForbidChatFormVisible = false;
        this.showSuccess();
        setTimeout(() => {
          this.loadJinYan();
        }, 300);
        // this.loadJinYan();
      });
    },
    loadJinYan() {
      this.listLoading = true;
      getJinYanPlayerList(this.jinyanListQuery).then(res => {
        this.jinyanUserList = res.itemArray;
        this.jinyanUserCount = res.total;
        this.listLoading = false;
      });
    },
    /** ********************禁默设置********************* */
    // 禁默设置
    handleJinMoPlatformChange(e) {
      const item = this.findPlatFormItem(e);
      if (item) {
        getCenterServerList(item.centerPlatformId).then(res => {
          this.jinMoServerList = res.itemArray;
        });
      }
    },
    handleJinMoServerChange(e) {
      const serverInfo = this.findJinMoServerItem(e);
      if (serverInfo) {
        this.jinMoListQuery.centerPlatformId = serverInfo.centerPlatformId;
        this.jinMoListQuery.centerServerId = serverInfo.serverId;
      }
    },
    findJinMoServerItem(serverId) {
      const server = this.jinMoServerList.find(n => {
        return n.id == serverId;
      });
      if (server) {
        return server;
      }
      return undefined;
    },

    handleJinMoFilter(e) {
      // 搜索
      if (
        !this.jinMoListQuery.centerPlatformId ||
        !this.jinMoListQuery.centerServerId
      ) {
        Message({
          message: "请选择服务器",
          type: "error",
          duration: 1.5 * 1000
        });
        return;
      }
      this.jinMoListQuery.pageIndex = 1;

      this.loadJinMo();
    },
    handleJinMoCurrentChange(e) {
      // 分页
      this.jinMoListQuery.pageIndex = e;
      this.loadJinMo();
    },
    handleJieJinMo(e) {
      // 解除禁言
      this.dialogUnJinMoChatFormVisible = true;
      this.unJinMoChatTemp = e;
      console.log(e);
    },
    updateUnJinMoChatPlayer(e) {
      const postdata = {
        centerPlatformId: this.unJinMoChatTemp.centerPlatformId,
        centerServerId: this.unJinMoChatTemp.centerServerId,
        playerId: this.unJinMoChatTemp.id
      };
      console.log(postdata);
      unIgnoreChatPlayer(postdata).then(res => {
        this.dialogUnJinMoChatFormVisible = false;
        this.showSuccess();
        setTimeout(() => {
          this.loadJinMo();
        }, 300);
      });
    },
    loadJinMo() {
      this.listLoading = true;
      getJinMoPlayerList(this.jinMoListQuery).then(res => {
        this.jinMoUserList = res.itemArray;
        this.jinMoUserCount = res.total;
        this.listLoading = false;
      });
    }
  }
};
</script>
