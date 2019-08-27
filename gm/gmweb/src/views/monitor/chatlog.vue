<template>
 <div>
    <div class="filter-container">
        <el-select v-model="listQuery.plId" placeholder="平台" style="width: 160px" class="filter-item" @change="handlePlatformChange">
          <el-option v-for="item in platformList" :key="item.platformId" :label="item.platformName" :value="item.platformId" />
        </el-select>
        <el-select v-model="listQuery.sid" collapse-tags placeholder="服务器" clearable style="width: 220px" class="filter-item" @change="handleServerChange">
          <el-option v-for="item in serverList" :key="item.id" :label="item.serverName" :value="item.id"/>
        </el-select>

        <el-input placeholder="玩家Id" v-model="listQuery.playerId" style="width: 200px;" class="filter-item"/>
        <el-select v-model="listQuery.chatTypeTmp" placeholder="频道" clearable style="width: 120px" class="filter-item">
              <el-option v-for="item in chatTypeList" :key="item.key" :label="item.name" :value="item.key"/>
            </el-select>
        <el-input placeholder="聊天内容" v-model="listQuery.content" style="width: 200px;" class="filter-item"/>
        <div class="filter-item">
            <el-date-picker v-model="listQuery.startEnd" type="datetimerange" range-separator="至" start-placeholder="开始时间" end-placeholder="结束时间">
            </el-date-picker>
        </div>
        <el-button v-waves class="filter-item" type="primary" icon="el-icon-search" @click="handleFilter">搜索</el-button>
    </div>
    <el-table
            v-loading="listLoading"
            :key="tableKey"
            :data="logData"
            border
            fit
            highlight-current-row
            style="width: 100%;margin-top:15px;">
            <el-table-column label="日志时间" align="center" width="160px">
                <template slot-scope="scope">
                    <span>{{ scope.row.logTime | parseTimeFilter }}</span>
                </template>
            </el-table-column>
            <el-table-column label="玩家名字" width="100px" align="left">
                <template slot-scope="scope">
                    <span>{{ scope.row.name}}</span>
                </template>
            </el-table-column>
            <el-table-column label="玩家角色" width="80px" align="left">
                <template slot-scope="scope">
                    <span>{{ scope.row.role}}</span>
                </template>
            </el-table-column>
            <el-table-column label="聊天内容" min-width="200px" align="left">
                <template slot-scope="scope">
                    <span>{{ scope.row.content | commonFilter("byte")}}</span>
                </template>
            </el-table-column>
             <el-table-column label="聊天频道" width="80px" align="left">
                <template slot-scope="scope">
                    <span>{{ scope.row.channel | parseChatTypeFilter}}</span>
                </template>
            </el-table-column>
            <el-table-column label="私聊目标id" width="120px" align="left">
                <template slot-scope="scope">
                    <span>{{ scope.row.recvId}}</span>
                </template>
            </el-table-column>
            <el-table-column label="私聊目标" width="120px" align="left">
                <template slot-scope="scope">
                    <span>{{ scope.row.recvName}}</span>
                </template>
            </el-table-column>
            <el-table-column label="玩家性别" width="80px" align="left">
                <template slot-scope="scope">
                    <span>{{ scope.row.sex}}</span>
                </template>
            </el-table-column>
            <el-table-column label="等级" width="50px" align="left">
                <template slot-scope="scope">
                    <span>{{ scope.row.level}}</span>
                </template>
            </el-table-column>

            <el-table-column label="VIP" width="50px" align="left">
                <template slot-scope="scope">
                    <span>{{ scope.row.vip}}</span>
                </template>
            </el-table-column>
            <el-table-column label="平台ID" width="80px" align="left">
                <template slot-scope="scope">
                    <span>{{ scope.row.platform}}</span>
                </template>
            </el-table-column>
            <el-table-column label="服务器ID" width="100px" align="left">
                <template slot-scope="scope">
                    <span>{{ scope.row.serverId}}</span>
                </template>
            </el-table-column>
            <el-table-column label="服务器类型" width="100px" align="left">
                <template slot-scope="scope">
                    <span>{{ scope.row.serverType}}</span>
                </template>
            </el-table-column>
            <el-table-column label="用户ID" width="80px" align="left">
                <template slot-scope="scope">
                    <span>{{ scope.row.userId}}</span>
                </template>
            </el-table-column>
            <el-table-column label="玩家ID" width="180px" align="left">
                <template slot-scope="scope">
                    <span>{{ scope.row.playerIdString}}</span>
                </template>
            </el-table-column>
            <el-table-column label="ip" width="100px" align="left">
                <template slot-scope="scope">
                    <span>{{ scope.row.ip}}</span>
                </template>
            </el-table-column>
            <el-table-column label="操作" fixed="right" align="center" width="100" class-name="small-padding fixed-width">
              <template slot-scope="scope">
                <el-button type="primary" size="mini" @click="handleForbid(scope.row)"
                v-permission="['super_admin', 'super_channel', 'channel', 'platform', 'service', 'minitor','super_channel_service','common_service','gaoji_service']">
                操作</el-button>
              </template>
          </el-table-column>
        </el-table>

        <div class="pagination-container" style="margin-top:15px;">
            <el-pagination :current-page="listQuery.pageIndex" :page-sizes="[20]" :total="totalCount" background layout="total, sizes, prev, pager, next, jumper"  @current-change="handleCurrentChange"/>
        </div>

        <el-dialog :visible.sync="dialogOperateFormVisible" title="聊天操作">
          <el-form ref="dataForm" :model="monitorChatTemp" label-position="left" label-width="120px" style="width: 400px; margin-left:50px;">
            <el-form-item label="玩家昵称">
              <el-input v-model="monitorChatTemp.name" :disabled="true"/>
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
            <el-button v-if="monitorChatTemp.forbid==0" type="primary" @click="updateForbidPlayerMonitor">封禁</el-button>
            <el-button v-if="monitorChatTemp.forbid==1" type="danger" @click="updateUnForbidPlayerMonitor">解禁</el-button>
            <el-button v-if="monitorChatTemp.forbidChat==0" type="primary" @click="updateForbidChatPlayerMonitor">禁言</el-button>
            <el-button v-if="monitorChatTemp.forbidChat==1" type="danger" @click="updateUnForbidChatPlayerMonitor">解言</el-button>
            <el-button v-if="monitorChatTemp.ignoreChat==0" type="primary" @click="updateIgnoreChatPlayerMonitor">禁默</el-button>
            <el-button v-if="monitorChatTemp.ignoreChat==1" type="danger" @click="updateUnIgnoreChatPlayerMonitor">解默</el-button>

            <el-button type="primary" @click="updateKickOutPlayer">踢人</el-button>
            <el-button v-if="monitorChatTemp.centerForbid==0" type="primary" @click="updateCenterUserForbid">中心封号</el-button>
            <el-button v-if="monitorChatTemp.centerForbid==1" type="danger" @click="updateUnCenterUserForbid">中心解封</el-button>
            <el-button v-if="monitorChatTemp.ipForbid==0" type="primary" @click="updateCenterIpForbid">中心封ip</el-button>
            <el-button v-if="monitorChatTemp.ipForbid==1" type="danger" @click="updateUnCenterIpForbid">解封ip</el-button>
          </div>
        </el-dialog>
 </div>    
    
</template>
<script>
import waves from "@/directive/waves"; // 水波纹指令
import permission from "@/directive/permission/index.js"; // 权限判断指令
import { getAllPlatformList } from "@/api/platform";
import { getCenterServerList } from "@/api/center";
import { getChatLog, getLogMeta, getLogMetaMsgList } from "@/api/log";
import { parseTime } from "@/utils/index";
import { binaryToStr, strToBinary } from "@/utils/binary";
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
import { chatTypeList, chatMethodList, chatForbidTimeList } from "@/types/chat";
import {
  getCenterUserInfo,
  updateCenterUserName,
  updateCenterForbid,
  updateCenterIpForbid,
  updateCenterIpUnForbid 
} from "@/api/centeruser";
export default {
  name: "PlayerChatLog",
  directives: {
    waves,
    permission
  },
  filters: {
    parseTimeFilter: function(value) {
      return parseTime(value, "{y}-{m}-{d} {h}:{i}:{s}");
    },
    parseChatTypeFilter: function(value) {
      let chatType = chatTypeList[value];
      if (chatType) {
        return chatType.name;
      }
      return "";
    },
    commonFilter: function(value, type) {
      if (type == "datetime") {
        return parseTime(value, "{y}-{m}-{d} {h}:{i}:{s}");
      }
      if (type == "byte") {
        return binaryToStr(value);
      }
      if (type == "normal") {
        return value;
      }
      return value;
    }
  },
  created() {
    this.initMetaData();
    // this.getList();
  },
  data() {
    return {
      tableKey: 1,
      //基础元数据
      metaLogType: [], //日志列表
      platformList: [], //平台列表
      serverList: [], //服务器列表
      listQuery: {
        tableName: "chat_content",
        beginTime: undefined,
        endTime: undefined,
        plId : undefined,
        platformId: undefined,
        sid: undefined,
        serverType: -1,
        startEnd: [],
        serverId: undefined,
        pageIndex: 1,
        playerId: undefined,
        chatContent: undefined,
        content: undefined
      },
      logData: [],
      totalCount: 0,
      listLoading: false,
      tableMetaMap: new Map(),
      metaColumnArray: [],
      monitorChatTemp: {},
      chatTypeList: [],
      chatForbidTimeArray: [],
      dialogOperateFormVisible: undefined
    };
  },
  methods: {
    handleLogTypeChange(e) {
      this.loadMetaColumn(e);
    },
    handlePlatformChange(e) {
      this.listQuery.sid = undefined;
      this.listQuery.serverType = -1;
      this.listQuery.serverid = undefined;
      if(this.listQuery.plId){
        let platform = this.findPlatformItem(e)
        if(platform){
          this.listQuery.platformId = platform.centerPlatformId;
        }
      }
      getCenterServerList(this.listQuery.platformId).then(res => {
        this.serverList = res.itemArray;
      });
      //   console.log(this.listQuery);
    },
    handleServerChange(e) {
      if (!e) {
        this.listQuery.serverType = -1;
        this.listQuery.serverid = undefined;
        return;
      }
      let item = this.findServerItem(e);
      if (item) {
        this.listQuery.serverId = item.serverId;
        this.listQuery.serverType = item.serverType;
      }
      //   console.log(this.listQuery);
    },
    handleForbid(e) {
      this.monitorChatTemp = Object.assign({}, e);
      const postData = {
        centerPlatformId: this.monitorChatTemp.platform,
        centerServerId: this.monitorChatTemp.serverId,
        playerId: this.monitorChatTemp.playerIdString
      };

      getPlayerInfo(postData).then(res => {
        this.monitorChatTemp.forbid = res.forbid;
        this.monitorChatTemp.forbidChat = res.forbidChat;
        this.monitorChatTemp.ignoreChat = res.ignoreChat;
        this.monitorChatTemp.centerForbid = res.centerForbid;
        this.monitorChatTemp.ipForbid = res.ipForbid;
        this.dialogOperateFormVisible = true;
      });
    },
    handleFilter(e) {
      console.log(this.listQuery.chatTypeTmp);
       if (this.listQuery.plId === undefined ||this.listQuery.plId === "") {
         this.$message.error("平台不能为空");
         return
      }
      if (
        this.listQuery.chatTypeTmp === undefined ||
        this.listQuery.chatTypeTmp === ""
      ) {
        console.log('here')
        this.listQuery.chatType = -1;
      }else{
        this.listQuery.chatType = parseInt(this.listQuery.chatTypeTmp)
      }
      this.listQuery.pageIndex = 1;
      if (this.listQuery.startEnd && this.listQuery.startEnd.length == 2) {
        this.listQuery.beginTime = this.listQuery.startEnd[0].valueOf();
        this.listQuery.endTime = this.listQuery.startEnd[1].valueOf();
      }
      // if (this.listQuery.content) {
      // this.listQuery.chatContent = strToBinary(this.listQuery.content)
      this.listQuery.chatContent = this.listQuery.content;
      // }
      this.loadData();
    },
    handleCurrentChange(e) {
      this.listQuery.pageIndex = e;
      this.loadData();
    },
    initMetaData() {
      this.chatTypeList = chatTypeList;
      this.chatForbidTimeArray = chatForbidTimeList;
      getLogMetaMsgList(1).then(res => {
        this.metaLogType = res;
        if (this.metaLogType && this.metaLogType.length > 0) {
          this.listQuery.tableName = this.metaLogType[0].key;
          this.loadMetaColumn(this.metaLogType[0].key);
        }
      });

      let startDate = new Date();
      startDate = new Date(
        startDate.getFullYear(),
        startDate.getMonth(),
        startDate.getDate()
      );
      let endDate = new Date();
      endDate.setDate(endDate.getDate() + 1);
      endDate = new Date(
        endDate.getFullYear(),
        endDate.getMonth(),
        endDate.getDate()
      );
      this.listQuery.startEnd = [startDate, endDate];
      getAllPlatformList().then(res => {
        this.platformList = res.itemArray;
      });
    },
    loadMetaColumn(e) {
      let metaData = this.tableMetaMap.get(e);
      if (metaData) {
        this.metaColumnArray = metaData;
        return;
      }
      getLogMeta(e, 1).then(res => {
        this.tableMetaMap.set(e, res);
        this.metaColumnArray = res;
      });
    },
    loadData() {
      if (!this.listQuery.tableName) {
        this.$message.error("日志类型不能为空");
        return;
      }
      this.listLoading = true;
      getChatLog(this.listQuery).then(res => {
        this.logData = res.itemArray;
        this.totalCount = res.totalCount;
        this.listLoading = false;
      });
    },
    updateForbidPlayerMonitor(e) {
      const postData = {
        centerPlatformId: this.monitorChatTemp.platform,
        centerServerId: this.monitorChatTemp.serverId,
        playerId: this.monitorChatTemp.playerIdString,
        reason: this.monitorChatTemp.reason,
        forbidTime: this.monitorChatTemp.forbidTime
      };

      console.log(postData);
      forbidPlayer(postData).then(res => {
        this.dialogOperateFormVisible = false;
        this.showSuccess();
      });
    },
    updateUnForbidPlayerMonitor(e) {
      const postdata = {
        centerPlatformId: this.monitorChatTemp.platform,
        centerServerId: this.monitorChatTemp.serverId,
        playerId: this.monitorChatTemp.playerIdString
      };
      unForbidPlayer(postdata).then(res => {
        this.dialogOperateFormVisible = false;
        this.showSuccess();
      });
    },
    updateForbidChatPlayerMonitor(e) {
      const postData = {
        centerPlatformId: this.monitorChatTemp.platform,
        centerServerId: this.monitorChatTemp.serverId,
        playerId: this.monitorChatTemp.playerIdString,
        reason: this.monitorChatTemp.reason,
        forbidTime: this.monitorChatTemp.forbidTime
      };
      console.log(postData);

      forbidChatPlayer(postData).then(res => {
        this.dialogOperateFormVisible = false;
        this.showSuccess();
      });
    },
    updateUnForbidChatPlayerMonitor(e) {
      const postdata = {
        centerPlatformId: this.monitorChatTemp.platform,
        centerServerId: this.monitorChatTemp.serverId,
        playerId: this.monitorChatTemp.playerIdString
      };
      unForbidChatPlayer(postdata).then(res => {
        this.dialogOperateFormVisible = false;
        this.showSuccess();
      });
    },
    updateIgnoreChatPlayerMonitor(e) {
      const postData = {
        centerPlatformId: this.monitorChatTemp.platform,
        centerServerId: this.monitorChatTemp.serverId,
        playerId: this.monitorChatTemp.playerIdString,
        reason: this.monitorChatTemp.reason,
        forbidTime: this.monitorChatTemp.forbidTime
      };
      console.log(postData);
      ignoreChatPlayer(postData).then(res => {
        this.dialogOperateFormVisible = false;
        this.showSuccess();
      });
    },
    updateUnIgnoreChatPlayerMonitor(e) {
      const postdata = {
        centerPlatformId: this.monitorChatTemp.platform,
        centerServerId: this.monitorChatTemp.serverId,
        playerId: this.monitorChatTemp.playerIdString
      };
      console.log(postdata);
      unIgnoreChatPlayer(postdata).then(res => {
        this.dialogOperateFormVisible = false;
        this.showSuccess();
      });
    },
    updateKickOutPlayer(e) {
      const postData = {
        centerPlatformId: this.monitorChatTemp.platform,
        centerServerId: this.monitorChatTemp.serverId,
        playerId: this.monitorChatTemp.playerIdString,
        reason: this.monitorChatTemp.reason
      };
      console.log(postData);
      kickOutPlayer(postData).then(res => {
        this.dialogOperateFormVisible = false;
        this.showSuccess();
      });
    },
    updateCenterUserForbid(e) {
      const postData = {
        userId: this.monitorChatTemp.userId,
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
    updateUnCenterUserForbid(e) {
      const postData = {
        userId: this.monitorChatTemp.userId,
        reason: this.monitorChatTemp.reason,
        forbid: 0,
        forbidTime: this.monitorChatTemp.forbidTime
      };
      console.log(postData);
      updateCenterForbid(postData).then(res => {
        this.dialogOperateFormVisible = false;
        this.showSuccess();
      });
    },
    updateCenterIpForbid(e) {
      const postData = {
        ip: this.monitorChatTemp.ip,
        reason: this.monitorChatTemp.reason,
        forbid: 1,
        forbidTime: this.monitorChatTemp.forbidTime,
        centerPlatformId: this.monitorChatTemp.platform,
        centerServerId: this.monitorChatTemp.serverId,
        playerId: this.monitorChatTemp.playerIdString,
      };
      console.log(postData);
      updateCenterIpForbid(postData).then(res => {
        this.dialogOperateFormVisible = false;
        this.showSuccess();
      });
    },
    updateUnCenterIpForbid(e) {
      const postData = {
        ip: this.monitorChatTemp.ip,
        reason: this.monitorChatTemp.reason,
        forbid: 0,
        forbidTime: this.monitorChatTemp.forbidTime
      };
      console.log(postData);
      updateCenterIpUnForbid(postData).then(res => {
        this.dialogOperateFormVisible = false;
        this.showSuccess();
      });
    },
    findServerItem(serverid) {
      const server = this.serverList.find(n => {
        return n.id == serverid;
      });
      if (server) {
        return server;
      }
      return undefined;
    },
    findPlatformItem(plId){
      const platform = this.platformList.find(n => {
        return n.platformId == plId;
      })
      if(platform){
        return platform
      }
      return undefined;
    },
    showSuccess() {
      this.$message({
        message: "设置成功",
        type: "success",
        duration: 1000
      });
    }
  }
};
</script>

