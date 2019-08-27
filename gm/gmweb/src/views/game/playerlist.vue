<template>
    <div class="app-container">
        <div class="filter-container">
            <el-select v-model="listQuery.channelId" placeholder="渠道" style="width: 120px" class="filter-item" @change="handleChannelChange">
                <el-option v-for="item in channelList" :key="item.channelId" :label="item.channelName" :value="item.channelId"/>
            </el-select>

            <el-select v-model="listQuery.platformId" placeholder="平台" style="width: 160px" class="filter-item" @change="handlePlatformChange">
                <el-option v-for="item in tempPlatformList" :key="item.platformId" :label="item.platformName" :value="item.platformId" />
            </el-select>

            <el-select v-model="listQuery.serverId" collapse-tags placeholder="服务器" clearable style="width: 220px" class="filter-item" >
              <el-option v-for="item in serverList" :key="item.id" :label="item.serverName" :value="item.id"/>
            </el-select>

            <el-input placeholder="玩家名" v-model="listQuery.playerName" style="width: 200px;" class="filter-item" @keyup.enter.native="handleFilter"/>

            <el-button v-waves class="filter-item" type="primary" icon="el-icon-search" @click="handleFilter">搜索</el-button>
        </div>

        <el-table
            v-loading="listLoading"
            :key="tableKey"
            :data="list"
            border
            fit
            highlight-current-row
            style="width: 100%;margin-top:15px;"
            @sort-change="handleSort">
            <el-table-column fixed="left" label="玩家Id" align="center" width="180px" sortable="custom" prop="1">
                <template slot-scope="scope">
                    <span>{{ scope.row.id }}</span>
                </template>
            </el-table-column>
            <el-table-column fixed="left" label="账户Id" min-width="150px" align="left" sortable="custom" prop="2">
                <template slot-scope="scope">
                    <span>{{ scope.row.userId}}</span>
                </template>
            </el-table-column>
            <el-table-column fixed="left" label="玩家名" min-width="150px" align="left" sortable="custom" prop="4">
                <template slot-scope="scope">
                    <span>{{ scope.row.name}}</span>
                </template>
            </el-table-column>
            <el-table-column label="SDK" min-width="110px" align="left" sortable="custom" prop="29">
                <template slot-scope="scope">
                    <span>{{ parseSdkType(scope.row.sdkType)}}</span>
                </template>
            </el-table-column>
            <el-table-column label="服务器" min-width="100px" align="left" sortable="custom" prop="3">
                <template slot-scope="scope">
                    <span>{{ scope.row.serverId}}</span>
                </template>
            </el-table-column>
            <el-table-column label="原始服务器" min-width="100px" align="left" sortable="custom" prop="3">
                <template slot-scope="scope">
                    <span>{{ scope.row.originServerId}}</span>
                </template>
            </el-table-column>
            <el-table-column label="角色" min-width="80px" align="left" sortable="custom" prop="5">
                <template slot-scope="scope">
                    <span>{{ scope.row.role | parsePlayerRole}}</span>
                </template>
            </el-table-column>
            <el-table-column label="性别" min-width="80px" align="left" sortable="custom" prop="6">
                <template slot-scope="scope">
                    <span>{{ scope.row.sex | parsesex}}</span>
                </template>
            </el-table-column>
            <el-table-column label="最后登陆时间" min-width="150px" align="left" sortable="custom" prop="7">
                <template slot-scope="scope">
                    <span>{{ scope.row.lastLoginTime | parseTime }}</span>
                </template>
            </el-table-column>
            <el-table-column label="最后登出时间" min-width="150px" align="left" sortable="custom" prop="8">
                <template slot-scope="scope">
                    <span>{{ scope.row.lastLogoutTime | parseTime}}</span>
                </template>
            </el-table-column>
            <el-table-column label="上线时长" min-width="150px" align="left" sortable="custom" prop="9">
                <template slot-scope="scope">
                    <span>{{ scope.row.onlineTime | parseSecond}}</span>
                </template>
            </el-table-column>
            <el-table-column label="等级" min-width="80px" align="left" sortable="custom" prop="14">
                <template slot-scope="scope">
                    <span>{{ scope.row.level}}</span>
                </template>
            </el-table-column>
            <el-table-column label="转数" min-width="80px" align="left" sortable="custom" prop="15">
                <template slot-scope="scope">
                    <span>{{ scope.row.zhuanSheng}}</span>
                </template>
            </el-table-column>
            <el-table-column label="银两" min-width="80px" align="left" sortable="custom" prop="16">
                <template slot-scope="scope">
                    <span>{{ scope.row.silver}}</span>
                </template>
            </el-table-column>
            <el-table-column label="元宝" min-width="80px" align="left" sortable="custom" prop="17">
                <template slot-scope="scope">
                    <span>{{ scope.row.gold}}</span>
                </template>
            </el-table-column>
            <el-table-column label="绑元" min-width="80px" align="left" sortable="custom" prop="18">
                <template slot-scope="scope">
                    <span>{{ scope.row.bindGold}}</span>
                </template>
            </el-table-column>
            <el-table-column label="原石" min-width="80px" align="left" sortable="custom" prop="19">
                <template slot-scope="scope">
                    <span>{{ scope.row.yuanshi}}</span>
                </template>
            </el-table-column>
            <el-table-column label="工会" min-width="150px" align="left" sortable="custom" prop="20">
                <template slot-scope="scope">
                    <span>{{ scope.row.allianceName}}</span>
                </template>
            </el-table-column>
            <el-table-column label="配偶" min-width="100px" align="left" sortable="custom" prop="21">
                <template slot-scope="scope">
                    <span>{{ scope.row.spouseName}}</span>
                </template>
            </el-table-column>
            <el-table-column label="魅力值" min-width="90px" align="left" sortable="custom" prop="22">
                <template slot-scope="scope">
                    <span>{{ scope.row.charm}}</span>
                </template>
            </el-table-column>
            <el-table-column label="战斗力" min-width="90px" align="left" sortable="custom" prop="23">
                <template slot-scope="scope">
                    <span>{{ scope.row.power}}</span>
                </template>
            </el-table-column>
            <el-table-column label="累计充值金额" min-width="90px" align="left" sortable="custom" prop="24">
                <template slot-scope="scope">
                    <span>{{ scope.row.totalChargeMoney}}</span>
                </template>
            </el-table-column>
            <el-table-column label="累计充值元宝" min-width="90px" align="left" sortable="custom" prop="27">
                <template slot-scope="scope">
                    <span>{{ scope.row.totalChargeGold}}</span>
                </template>
            </el-table-column>
            <el-table-column label="累计扶持元宝" min-width="90px" align="left" sortable="custom" prop="28">
                <template slot-scope="scope">
                    <span>{{ scope.row.totalPrivilegeChargeGold}}</span>
                </template>
            </el-table-column>
            <el-table-column label="昨日充值总金额" min-width="90px" align="left" sortable="custom" prop="25">
                <template slot-scope="scope">
                    <span>{{ scope.row.yesterdayChargeMoney}}</span>
                </template>
            </el-table-column>
            <el-table-column label="今日充值总金额" min-width="90px" align="left" sortable="custom" prop="26">
                <template slot-scope="scope">
                    <span>{{ scope.row.todayChargeMoney}}</span>
                </template>
            </el-table-column>
            <!-- <el-table-column label="下线时长" min-width="150px" align="left" sortable="custom" prop="10">
                <template slot-scope="scope">
                    <span>{{ scope.row.offlineTime}}</span>
                </template>
            </el-table-column> -->
            <el-table-column label="总在线时长" min-width="150px" align="left" sortable="custom" prop="11">
                <template slot-scope="scope">
                    <span>{{ scope.row.totalOnlineTime | parseSecond}}</span>
                </template>
            </el-table-column>
            <el-table-column label="当日在线时长" min-width="150px" align="left" sortable="custom" prop="12">
                <template slot-scope="scope">
                    <span>{{ scope.row.todayOnlineTime | parseSecond}}</span>
                </template>
            </el-table-column>
            <el-table-column label="创建时间" min-width="150px" align="left" sortable="custom" prop="13">
                <template slot-scope="scope">
                    <span>{{ scope.row.createTime | parseTime}}</span>
                </template>
            </el-table-column>
            <el-table-column fixed="right" label="操作" align="center" width="320" class-name="small-padding fixed-width">
                <template slot-scope="scope">
                <!-- <el-button type="primary" v-if="scope.row.forbid == 1" size="mini" @click="handleUnForbit(scope.row)">解禁</el-button>
                <el-button size="mini" type="danger" v-if="scope.row.forbid != 1" @click="handleForbit(scope.row)">封禁</el-button> -->
                <el-button type="primary" size="mini" @click="handleForbid(scope.row)">封禁</el-button>
                <el-button type="primary" size="mini" @click="handleKickOut(scope.row)">踢人</el-button>
                <el-button type="primary" size="mini" @click="handleView(scope.row)">查看</el-button>
                <el-button type="primary" size="mini" @click="handleSet(scope.row)">设置</el-button>
                </template>
            </el-table-column>
        </el-table>

        <div class="pagination-container" style="margin-top:15px;">
            <el-pagination :current-page="listQuery.pageIndex" :page-sizes="[20]" :total="total" background layout="total, sizes, prev, pager, next, jumper"  @current-change="handleCurrentChange"/>
        </div>

        <el-dialog :visible.sync="dialogForbidFormVisible" title="封禁用户">
          <el-form ref="dataForm" :model="monitorTemp" label-position="left" label-width="120px" style="width: 400px; margin-left:50px;">
            <el-form-item label="封禁用户名">
              <el-input v-model="monitorTemp.name" :disabled="true"/>
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

        <el-dialog :visible.sync="dialogUnForbidFormVisible" title="是否解禁用户">
          <el-form ref="dataForm" :model="monitorTemp" label-position="left" label-width="120px" style="width: 400px; margin-left:50px;">
            <el-form-item label="封禁用户名">
              <el-input v-model="monitorTemp.name" :disabled="true"/>
            </el-form-item>
          </el-form>
          <div slot="footer" class="dialog-footer">
            <el-button @click="dialogUnForbidFormVisible = false">取消</el-button>
            <el-button type="primary" @click="updateUnForbidPlayer">解禁</el-button>
          </div>
        </el-dialog>
        <el-dialog :visible.sync="dialogOperateFormVisible" title="玩家账户操作">
          <el-form ref="dataForm" :model="monitorChatTemp" label-position="left" label-width="120px" style="width: 400px; margin-left:50px;">
            <el-form-item label="玩家昵称">
              <el-input v-model="monitorChatTemp.name" :disabled="true"/>
            </el-form-item>
            <el-form-item label="IP">
              <el-input v-model="monitorChatTemp.ip" :disabled="true"/>
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

            <el-button v-if="monitorChatTemp.centerForbid==0" type="primary" @click="updateCenterUserForbid">中心封号</el-button>
            <el-button v-if="monitorChatTemp.centerForbid==1" type="danger" @click="updateUnCenterUserForbid">中心解封</el-button>
            <el-button v-if="monitorChatTemp.ipForbid==0" type="primary" @click="updateCenterIpForbid">中心封ip</el-button>
            <el-button v-if="monitorChatTemp.ipForbid==1" type="danger" @click="updateUnCenterIpForbid">解封ip</el-button>
          </div>
        </el-dialog>
        <el-dialog :visible.sync="dialogUserNameVisible" title="设置用户密码" width="30%">
            <el-form ref="dataForm" :model="userNameTemp" label-position="left" label-width="100px" style="width: 300px; margin-left:50px;">
                <el-form-item label="用户id">
                    <el-input v-model="userNameTemp.userId" :disabled="true"/>
                </el-form-item>
                <el-form-item label="用户名">
                    <el-input v-model="userNameTemp.name"/>
                </el-form-item>
                <el-form-item label="密码">
                    <el-input v-model="userNameTemp.password" type="password"/>
                </el-form-item>
            </el-form>
            <span slot="footer" class="dialog-footer">
                <el-button @click="dialogUserNameVisible = false">取消</el-button>
                <el-button type="primary" @click="updateUserName">确定</el-button>
            </span>
        </el-dialog>
        <el-dialog :visible.sync="dialogKickoutVisible" title="是否踢出用户">
          <el-form ref="dataForm" :model="monitorChatTemp" label-position="left" label-width="120px" style="width: 400px; margin-left:50px;">
            <el-form-item label="踢出用户名">
              <el-input v-model="monitorChatTemp.name" :disabled="true"/>
            </el-form-item>
            <el-form-item label="踢出原因">
              <el-input v-model="monitorChatTemp.reason"/>
            </el-form-item>
          </el-form>
          <div slot="footer" class="dialog-footer">
            <el-button @click="dialogKickoutVisible = false">取消</el-button>
            <el-button type="primary" @click="updateKickOutPlayer">踢出</el-button>
          </div>
        </el-dialog>
    </div>
</template>

<script>
import waves from "@/directive/waves"; // 水波纹指令
import permission from "@/directive/permission/index.js"; // 权限判断指令
import { getPlayerList } from "@/api/player";
import { parseTime } from "@/utils/index";
import { getAllChannel } from "@/api/channel";
import { getAllPlatformList } from "@/api/platform";
import { getCenterServerList,getAllSdkType } from "@/api/center";
import { Message, MessageBox } from "element-ui";
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
import { chatForbidTimeList } from "@/types/chat";
import { playerRoleMap } from "@/types/player";
import {
  getCenterUserInfo,
  updateCenterUserName,
  updateCenterForbid,
  updateCenterIpForbid,
  updateCenterIpUnForbid 
} from "@/api/centeruser";

export default {
  name: "PlayerList",
  directives: {
    waves,
    permission
  },
  filters: {
    parseTime: function(value) {
      return parseTime(value, "{y}-{m}-{d} {h}:{i}");
    },
    parseSecond: function(value) {
      let hour = parseInt(value / 60 / 60 / 1000);
      let reseMinute = value % (60 * 60);
      let minute = parseInt(reseMinute / 60);
      let reseSecond = reseMinute % 60;
      return hour + "时" + minute + "分" + reseSecond + "秒";
    },
    parsesex: function(value) {
      if (value == 1) {
        return "男";
      }
      if (value == 2) {
        return "女";
      }
      return value;
    },
    parsePlayerRole: function(value) {
      let info = playerRoleMap[value - 1];
      if (info) {
        return info.name;
      }
      return "";
    }
  },
  created() {
    this.initMetaData();
    // this.getList();
  },
  data() {
    return {
      listLoading: false,
      tableKey: 0,
      total: 0,
      listQuery: {
        pageIndex: 1,
        playerName: "",
        ordercol: 1,
        ordertype: 0,
        platformId: undefined,
        channelId: undefined,
        serverId: undefined
      },
      textMap: {
        update: "编辑",
        create: "添加"
      },
      dialogStatus: "",
      dialogPvVisible: false,
      dialogFormVisible: false,
      dialogUserNameVisible: false,
      temp: {},
      userNameTemp: {},
      list: [],
      dialogUnForbidFormVisible: false,
      dialogForbidFormVisible: false,
      dialogOperateFormVisible: false,
      dialogKickoutVisible: false,
      channelList: [],
      platformList: [],
      groupList: [],
      tempPlatformList: [],
      serverList: [],
      chatForbidTimeArray: [],
      monitorTemp: {},
      realTemp: {},
      monitorChatTemp: {},
      sdkList:[]
    };
  },
  methods: {
    handleFilter: function() {
      if (!this.listQuery.serverId) {
        Message({
          message: "请选择服务器",
          type: "error",
          duration: 1.5 * 1000
        });
        return;
      }
      console.log(this.listQuery);
      this.listQuery.pageIndex = 1;
      this.listQuery.ordercol = 1;
      this.listQuery.ordertype = 0;
      this.getList();
    },

    getList() {
      this.listLoading = true;
      getPlayerList(this.listQuery)
        .then(res => {
          this.list = res.itemArray;
          this.total = res.total;
          this.listLoading = false;
        })
        .catch(() => {
          this.listLoading = false;
        });
    },
    handleCurrentChange(e) {
      console.log(e);
      this.listQuery.pageIndex = e;
      this.getList();
    },
    handleSort(e) {
      if (!this.listQuery.serverId) {
        Message({
          message: "请选择服务器",
          type: "error",
          duration: 1.5 * 1000
        });
        return;
      }
      this.listQuery.ordercol = parseInt(e.prop);
      this.listQuery.ordertype = 0;
      if (e.order == "descending") {
        this.listQuery.ordertype = 1;
      }
      this.getList();
    },
    initMetaData() {
      this.chatForbidTimeArray = chatForbidTimeList;
      getAllChannel().then(res => {
        this.channelList = res.itemArray;
      });
      getAllPlatformList().then(res => {
        this.platformList = res.itemArray;
        // this.tempPlatformList = this.platformList;
      });
      getAllSdkType().then(res =>{
        this.sdkList = res.itemArray;
      });
    },
    handleChannelChange: function(e) {
      if (e) {
        this.listQuery.platformId = undefined;
        this.tempPlatformList = this.findPlatFormList(e);
        if (this.tempPlatFormList && this.tempPlatFormList.length > 0) {
          this.listQuery.platformId = this.tempPlatFormList[0].platformId;
        }
        this.groupList = [];
        this.listQuery.serverId = undefined;
      }
    },
    handlePlatformChange: function(e) {
      console.log(e);
      if (e) {
        let item = this.findPlatFormItem(e);

        if (item) {
          getCenterServerList(item.centerPlatformId).then(res => {
            this.listQuery.serverId = undefined;
            this.serverList = res.itemArray;
          });
        }
      }
    },
    findPlatFormList(channelId) {
      if (!this.platformList || this.platformList.length == 0) {
        return;
      }
      return this.platformList.filter(function(item, index) {
        return item.channelId == channelId;
      });
    },
    findPlatFormItem(platformId) {
      let platform = this.platformList.find(n => {
        return n.platformId == platformId;
      });
      if (platform) {
        return platform;
      }
      return undefined;
    },
    findServerItem(id) {
      let server = this.serverList.find(n => {
        return n.id == id;
      });
      if (server) {
        return server;
      }
      return undefined;
    },
    handleUnForbit(e) {
      this.dialogUnForbidFormVisible = true;
      this.monitorTemp = e;
      this.monitorTemp.reason = undefined;
    },
    handleForbit(e) {
      this.dialogForbidFormVisible = true;
      this.monitorTemp = e;
      this.monitorTemp.reason = undefined;
    },
    updateUnForbidPlayer(e) {
      let myserver = this.findServerItem(this.listQuery.serverId);
      if (!myserver) {
        return;
      }
      const postdata = {
        centerPlatformId: myserver.centerPlatformId,
        centerServerId: myserver.serverId,
        playerId: this.monitorTemp.id
      };
      unForbidPlayer(postdata).then(res => {
        this.dialogUnForbidFormVisible = false;
        this.showSuccess();
        this.monitorTemp.forbid = 0;
        // this.getList();
      });
    },
    updateForbidPlayer(e) {
      let myserver = this.findServerItem(this.listQuery.serverId);
      if (!myserver) {
        return;
      }
      const postData = {
        centerPlatformId: myserver.centerPlatformId,
        centerServerId: myserver.serverId,
        playerId: this.monitorTemp.id,
        reason: this.monitorTemp.reason,
        forbidTime: this.monitorTemp.forbidTime
      };

      console.log(postData);
      forbidPlayer(postData).then(res => {
        this.dialogForbidFormVisible = false;
        this.showSuccess();
        // this.getList();
        this.monitorTemp.forbid = 1;
      });
    },
    handleView(e) {
      this.$router.push({
        name: "playerInfo",
        params: { id: e.id, serverId: this.listQuery.serverId }
      });
    },
    handleSet(e) {
      this.userNameTemp = Object.assign({}, e);
      this.userNameTemp.id = this.userNameTemp.userId;
      getCenterUserInfo(this.userNameTemp).then(res => {
        this.userNameTemp.name = res.name;
        this.dialogUserNameVisible = true;
      });
    },
    //玩家账户操作内容
    handleForbid(e) {
      this.monitorChatTemp = Object.assign({}, e);
      let platformInfo = this.findPlatFormItem(this.listQuery.platformId);
      if (!platformInfo) {
        return;
      }
      this.monitorChatTemp.platform = platformInfo.centerPlatformId;
      const postData = {
        centerPlatformId: platformInfo.centerPlatformId,
        centerServerId: this.monitorChatTemp.serverId,
        playerId: this.monitorChatTemp.id,
        ip: this.monitorChatTemp.ip
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
    updateForbidPlayerMonitor(e) {
      const postData = {
        centerPlatformId: this.monitorChatTemp.platform,
        centerServerId: this.monitorChatTemp.serverId,
        playerId: this.monitorChatTemp.id,
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
        playerId: this.monitorChatTemp.id
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
        playerId: this.monitorChatTemp.id,
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
        playerId: this.monitorChatTemp.id
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
        playerId: this.monitorChatTemp.id,
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
        playerId: this.monitorChatTemp.id
      };
      console.log(postdata);
      unIgnoreChatPlayer(postdata).then(res => {
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
        playerId: this.monitorChatTemp.id
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
    handleKickOut(e) {
      this.monitorChatTemp = Object.assign({}, e);
      let platformInfo = this.findPlatFormItem(this.listQuery.platformId);
      if (!platformInfo) {
        return;
      }
      this.monitorChatTemp.platform = platformInfo.centerPlatformId;
      this.dialogKickoutVisible = true;
    },
    updateKickOutPlayer(e) {
      const postData = {
        centerPlatformId: this.monitorChatTemp.platform,
        centerServerId: this.monitorChatTemp.serverId,
        playerId: this.monitorChatTemp.id,
        reason: this.monitorChatTemp.reason
      };
      console.log(postData);
      kickOutPlayer(postData).then(res => {
        this.dialogKickoutVisible = false;
        this.showSuccess();
      });
    },
    updateUserName: function(e) {
      updateCenterUserName(this.userNameTemp).then(res => {
        this.dialogUserNameVisible = false;
        this.showSuccess();
        // this.getList();
      });
    },
    parseSdkType:function(e){
      for(let i=0;i<this.sdkList.length;i++){
        let item = this.sdkList[i];
        if(item.key == e){
          return item.name;
        }
      }
      return e;
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

