<template>
    <div>
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
            <el-button class="filter-item" style="margin-left: 10px;" type="primary" icon="el-icon-edit" @click="handleFuChi">添加扶持号</el-button>
            <el-button class="filter-item" style="margin-left: 10px;" type="primary" icon="el-icon-edit" @click="handlePiLiangFuChi">批量扶持元宝</el-button>
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
                    <span v-if="scope.row.totalChargeMoney > 0"><font color="#DC143C" ><b>{{ scope.row.name}}</b></font></span>
                    <span v-if="scope.row.totalChargeMoney <= 0" >{{ scope.row.name}}</span>
                </template>
            </el-table-column>
              <el-table-column label="原始服务器" min-width="100px" align="left" sortable="custom" prop="3">
                <template slot-scope="scope">
                    <span>{{ scope.row.originServerId}}</span>
                </template>
            </el-table-column>
            <el-table-column label="累计充值金额" min-width="90px" align="left" sortable="custom" prop="24">
                <template slot-scope="scope">
                    <span>{{ scope.row.totalChargeMoney}}</span>
                </template>
            </el-table-column>
            <el-table-column label="累计充值元宝" min-width="90px" align="left" sortable="custom" prop="25">
                <template slot-scope="scope">
                    <span>{{ scope.row.totalChargeGold}}</span>
                </template>
            </el-table-column>
            <el-table-column label="累计扶持元宝" min-width="90px" align="left" sortable="custom" prop="26">
                <template slot-scope="scope">
                    <span>{{ scope.row.totalPrivilegeChargeGold}}</span>
                </template>
            </el-table-column>
            <el-table-column label="等级" min-width="80px" align="left" sortable="custom" prop="14">
                <template slot-scope="scope">
                    <span>{{ scope.row.level}}</span>
                </template>
            </el-table-column>
            <el-table-column label="最后登陆时间" min-width="150px" align="left" sortable="custom" prop="7">
                <template slot-scope="scope">
                    <span>{{ scope.row.lastLoginTime | parseTime }}</span>
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
            
            <!-- <el-table-column label="配偶" min-width="100px" align="left" sortable="custom" prop="21">
                <template slot-scope="scope">
                    <span>{{ scope.row.spouseName}}</span>
                </template>
            </el-table-column> -->
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
             <el-table-column label="扶持类型" min-width="120px" align="left" sortable="custom" prop="23">
                <template slot-scope="scope">
                    <span>{{ scope.row.privilegeType | parseFuchiType}}</span>
                </template>
            </el-table-column>
            <el-table-column fixed="right" label="操作" align="center" width="200" class-name="small-padding fixed-width">
                <template slot-scope="scope">
                <el-button type="primary" size="mini" @click="handleUnForbit(scope.row)">扶持</el-button>
                <el-button type="primary" size="mini" @click="handleUpdate(scope.row)">编辑</el-button>
                </template>
            </el-table-column>
        </el-table>

        <div class="pagination-container" style="margin-top:15px;">
            <el-pagination :current-page="listQuery.pageIndex" :page-sizes="[20]" :total="total" background layout="total, sizes, prev, pager, next, jumper"  @current-change="handleCurrentChange"/>
        </div>

        <el-dialog :visible.sync="dialogUnForbidFormVisible" title="扶持元宝">
          <el-form ref="dataForm" :model="monitorTemp" label-position="left" label-width="120px" style="width: 400px; margin-left:50px;">
            <el-form-item label="角色ID">
              <el-input v-model="monitorTemp.playerId" :disabled="true" />
            </el-form-item>
            <el-form-item label="扶持数量">
              <el-select v-model="monitorTemp.gold" placeholder="扶持数量" style="width: 160px" class="filter-item">
                <el-option v-for="item in supportAmountArray" :key="item.key" :label="item.name" :value="item.key" />
              </el-select>
            </el-form-item>
            <el-form-item label="扶持次数">
              <el-input-number v-model="monitorTemp.num" label="扶持次数"></el-input-number>
            </el-form-item>
            <el-form-item label="扶持原因">
              <el-input v-model="monitorTemp.reason"/>
            </el-form-item>
          </el-form>
          <div slot="footer" class="dialog-footer">
            <el-button @click="dialogUnForbidFormVisible = false">取消</el-button>
            <el-button type="primary" @click="updateFuchiPlayerGold">扶持</el-button>
          </div>
        </el-dialog>

        <el-dialog :visible.sync="dialogFuChiFormVisible" title="扶持用户(角色id或角色名二选一，都填以角色id优先，多个以英文逗号隔开)">
          <el-form ref="dataForm" :model="fuchiInfo" label-position="left" label-width="120px" style="width: 400px; margin-left:50px;">
            <el-form-item label="角色ID">
              <el-input v-model="fuchiInfo.playerId" :disabled="!(dialogStatus=='create')" />
            </el-form-item>
            <el-form-item label="角色名">
              <el-input v-model="fuchiInfo.name" :disabled="!(dialogStatus=='create')" />
            </el-form-item>
            <el-form-item label="扶持类型">
              <el-select v-model="fuchiInfo.privilege" placeholder="扶持类型" style="width: 160px" class="filter-item">
                  <el-option v-for="item in playerSupportList" :key="item.key" :label="item.name" :value="item.key" />
              </el-select>
            </el-form-item>
          </el-form>
          <div slot="footer" class="dialog-footer">
            <el-button @click="dialogFuChiFormVisible = false">取消</el-button>
            <el-button v-if="dialogStatus=='create'" type="primary" @click="fuChiPlayer">创建</el-button>
                <el-button v-else type="primary" @click="fuChiPlayer">确定</el-button>
          </div>
        </el-dialog>

        <el-dialog :visible.sync="dialogPiLiangFuChiFormVisible" title="批量扶持元宝(角色id或角色名二选一，都填以角色id优先，多个以英文逗号隔开)">
          <el-form ref="dataForm" :model="piFuchiInfo" label-position="left" label-width="120px" style="width: 400px; margin-left:50px;">
            <el-form-item label="发放方式">
              <template>
                <el-radio v-model="piFuchiInfo.allServer" :label="false">部分发放</el-radio>
                <el-radio v-model="piFuchiInfo.allServer" :label="true">全服发放</el-radio>
              </template>
            </el-form-item>
            <el-form-item label="角色ID">
              <el-input v-model="piFuchiInfo.playerId" :disabled="piFuchiInfo.allServer" />
            </el-form-item>
            <el-form-item label="角色名">
              <el-input v-model="piFuchiInfo.name" :disabled="piFuchiInfo.allServer" />
            </el-form-item>
            <el-form-item label="扶持数量">
              <el-select v-model="piFuchiInfo.gold" placeholder="扶持数量" style="width: 160px" class="filter-item">
                <el-option v-for="item in supportAmountArray" :key="item.key" :label="item.name" :value="item.key" />
              </el-select>
            </el-form-item>
            <el-form-item label="扶持次数">
              <el-input-number v-model="piFuchiInfo.num" label="扶持次数"></el-input-number>
            </el-form-item>
            <el-form-item label="扶持原因">
              <el-input v-model="piFuchiInfo.reason"/>
            </el-form-item>
          </el-form>
          <div slot="footer" class="dialog-footer">
              <el-button @click="dialogPiLiangFuChiFormVisible = false">取消</el-button>
              <el-button type="primary" @click="piLiangFuChiPlayer">确定</el-button>
          </div>
        </el-dialog>
    </div>
</template>

<script>
import waves from "@/directive/waves"; // 水波纹指令
import {
  getPlayerList,
  privilegeCharge,
  privilegeSet,
  privilegeChargeMulity
} from "@/api/supportplayer";
import { parseTime } from "@/utils/index";
import { getAllChannel } from "@/api/channel";
import { getAllPlatformList } from "@/api/platform";
import { getCenterServerList } from "@/api/center";
import { Message, MessageBox } from "element-ui";
import { chatForbidTimeList } from "@/types/chat";
import { playerSupportMap } from "@/types/player";
import { supportAmount } from "@/types/manage";

export default {
  name: "SupportPlayer",
  directives: {
    waves
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
    parseFuchiType: function(value) {
      let item = playerSupportMap[value];
      if (item) {
        return item.name;
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
      temp: {},
      list: [],
      dialogUnForbidFormVisible: false,
      dialogFuChiFormVisible: false,
      dialogPiLiangFuChiFormVisible: false,
      channelList: [],
      platformList: [],
      groupList: [],
      tempPlatformList: [],
      serverList: [],
      chatForbidTimeArray: [],
      monitorTemp: {
        channelId: undefined,
        platformId: undefined,
        serverId: undefined,
        playerId: undefined,
        gold: undefined,
        reason: undefined,
        num:1
      },
      playerSupportList: [],
      fuchiInfo: {
        playerId: undefined,
        gold: undefined,
        privilege: 1
      },
      piFuchiInfo: {
        playerId: undefined,
        name: undefined,
        gold: undefined,
        reason: undefined,
        allServer: false,
        num:1
      },
      supportAmountArray: []
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
      this.listQuery.ordercol = 24;
      this.listQuery.ordertype = 1;
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
      this.playerSupportList = playerSupportMap;
      this.supportAmountArray = supportAmount;
      getAllChannel().then(res => {
        this.channelList = res.itemArray;
      });
      getAllPlatformList().then(res => {
        this.platformList = res.itemArray;
        // this.tempPlatformList = this.platformList;
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
      // this.monitorTemp = e;
      this.monitorTemp.serverId = this.listQuery.serverId;
      this.monitorTemp.playerId = e.id;
      this.monitorTemp.reason = undefined;
      this.monitorTemp.channelId = this.listQuery.channelId;
      this.monitorTemp.platformId = this.listQuery.platformId;
      this.monitorTemp.playerName = e.name;
      this.monitorTemp.num = 1;
    },
    handleFuChi(e) {
      if (!this.listQuery.serverId) {
        Message({
          message: "请选择服务器",
          type: "error",
          duration: 1.5 * 1000
        });
        return;
      }
      this.dialogStatus = "create";
      this.dialogFuChiFormVisible = true;
      this.fuchiInfo = {
        playerId: undefined,
        gold: undefined,
        privilege: 1
      };
    },
    handlePiLiangFuChi(e) {
      if (!this.listQuery.serverId) {
        Message({
          message: "请选择服务器",
          type: "error",
          duration: 1.5 * 1000
        });
        return;
      }
      this.dialogPiLiangFuChiFormVisible = true;
      this.piFuchiInfo = {
        playerId: undefined,
        name: undefined,
        gold: undefined,
        reason: undefined,
        allServer: false,
        num:1
      };
    },
    handleUpdate(e) {
      this.fuchiInfo = {
        playerId: e.id,
        privilege: e.privilegeType,
        name: e.name
      };
      this.dialogStatus = "update";
      this.dialogFuChiFormVisible = true;
    },
    updateFuchiPlayerGold(e) {
      if (!this.listQuery.serverId) {
        Message({
          message: "请选择服务器",
          type: "error",
          duration: 1.5 * 1000
        });
        return;
      }

      const postData = {
        serverId: this.listQuery.serverId,
        playerId: this.monitorTemp.playerId.toString(),
        gold: this.monitorTemp.gold,
        reason: this.monitorTemp.reason,
        channelId: this.monitorTemp.channelId,
        platformId: this.monitorTemp.platformId,
        playerName: this.monitorTemp.playerName,
        num: this.monitorTemp.num
      };

      console.log(postData);
      privilegeCharge(postData).then(res => {
        this.dialogUnForbidFormVisible = false;
        this.showSuccess();
        setTimeout(() => {
          this.getList();
        }, 500);
      });
    },
    fuChiPlayer(e) {
      if (!this.listQuery.serverId) {
        Message({
          message: "请选择服务器",
          type: "error",
          duration: 1.5 * 1000
        });
        return;
      }
      const postData = {
        serverId: this.listQuery.serverId,
        playerId: this.fuchiInfo.playerId,
        privilege: this.fuchiInfo.privilege,
        playerName: this.fuchiInfo.name
      };

      console.log(postData);
      privilegeSet(postData).then(res => {
        this.dialogFuChiFormVisible = false;
        this.showSuccess();
        setTimeout(() => {
          this.getList();
        }, 1000);
        // this.getList();
      });
    },
    piLiangFuChiPlayer() {
      if (!this.listQuery.serverId) {
        Message({
          message: "请选择服务器",
          type: "error",
          duration: 1.5 * 1000
        });
        return;
      }

      const postData = {
        serverId: this.listQuery.serverId,
        playerId: this.piFuchiInfo.playerId,
        playerName: this.piFuchiInfo.name,
        gold: this.piFuchiInfo.gold,
        reason: this.piFuchiInfo.reason,
        channelId: this.listQuery.channelId,
        platformId: this.listQuery.platformId,
        allServer: this.piFuchiInfo.allServer,
        num: this.piFuchiInfo.num
      };

      console.log(postData);
      privilegeChargeMulity(postData).then(res => {
        this.dialogPiLiangFuChiFormVisible = false;
        this.showSuccess();
        setTimeout(() => {
          this.getList();
        }, 500);
      });
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

