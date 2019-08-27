<template>
    <div class="app-container">
        <div class="filter-container">
            <el-input placeholder="邮件标题" v-model="listQuery.title" style="width: 200px;" class="filter-item"/>
            <el-select v-model="listQuery.mailState" class="filter-item" clearable placeholder="服务器状态">
                <el-option v-for="item in mailStateList" :key="item.key" :label="item.name" :value="item.key"/>
            </el-select>
            <el-button v-waves class="filter-item" type="primary" icon="el-icon-search" @click="handleFilter">搜索</el-button>
            <el-button class="filter-item" style="margin-left: 10px;" type="primary" icon="el-icon-edit" @click="handleCreate">添加</el-button>
        </div>

        <el-table
            v-loading="listLoading"
            :key="tableKey"
            :data="list"
            border
            fit
            highlight-current-row
            style="width: 100%;margin-top:15px;">
            <el-table-column  label="邮件Id" align="center" width="100px" >
                <template slot-scope="scope">
                    <span>{{ scope.row.id }}</span>
                </template>
            </el-table-column>
            <el-table-column  label="邮件类型" min-width="80px" align="left" >
                <template slot-scope="scope">
                    <span>{{ scope.row.mailType | parseMailType}}</span>
                </template>
            </el-table-column>
            <el-table-column  label="服务器名" min-width="150px" align="left" >
                <template slot-scope="scope">
                    <span>{{ scope.row.serverName}}</span>
                </template>
            </el-table-column>
            <el-table-column  label="标题" min-width="150px" align="left" >
                <template slot-scope="scope">
                    <span>{{ scope.row.title}}</span>
                </template>
            </el-table-column>
            <el-table-column  label="内容" min-width="200px" align="left" >
                <template slot-scope="scope">
                    <span>{{ scope.row.content}}</span>
                </template>
            </el-table-column>
            <!-- <el-table-column  label="冻结时间(分)" min-width="150px" align="left" >
                <template slot-scope="scope">
                    <span>{{ scope.row.freezTime}}</span>
                </template>
            </el-table-column>
            <el-table-column  label="有效天数" min-width="80px" align="left" >
                <template slot-scope="scope">
                    <span>{{ scope.row.effectDays}}</span>
                </template>
            </el-table-column> -->
            <el-table-column  label="角色限制时间" min-width="150px" align="left" >
                <template slot-scope="scope">
                    <span>{{ scope.row.roleStartTime | parseTime}}</span>
                </template>
            </el-table-column>
            <!-- <el-table-column  label="角色限制结束时间" min-width="150px" align="left" >
                <template slot-scope="scope">
                    <span>{{ scope.row.roleEndTime| parseTime}}</span>
                </template>
            </el-table-column> -->
            <el-table-column  label="限制等级" min-width="100px" align="left" >
                <template slot-scope="scope">
                    <span>{{ scope.row.minLevel}}</span>
                </template>
            </el-table-column>
            <el-table-column  label="是否绑定" min-width="100px" align="left" >
                <template slot-scope="scope">
                    <span>{{ scope.row.bindFlag | parseYesOrNo}}</span>
                </template>
            </el-table-column>
            <!-- <el-table-column  label="限制最大等级" min-width="100px" align="left" >
                <template slot-scope="scope">
                    <span>{{ scope.row.maxLevel}}</span>
                </template>
            </el-table-column> -->
            <el-table-column  label="邮件状态" min-width="80px" align="left" >
                <template slot-scope="scope">
                    <span>{{ scope.row.mailState | parseMailState}}</span>
                </template>
            </el-table-column>
            <el-table-column  label="备注" min-width="80px" align="left" >
                <template slot-scope="scope">
                    <span>{{ scope.row.remark}}</span>
                </template>
            </el-table-column>
            <el-table-column  label="审核理由" min-width="80px" align="left" >
                <template slot-scope="scope">
                    <span>{{ scope.row.approveReason}}</span>
                </template>
            </el-table-column>
            <!-- <el-table-column  label="删除状态" width="80px" align="left" >
                <template slot-scope="scope">
                    <span>{{ scope.row.deleteTime > 0 ? '已删除' : ''}}</span>
                </template>
            </el-table-column> -->
            <el-table-column  label="创建时间" min-width="150px" align="left" >
                <template slot-scope="scope">
                    <span>{{ scope.row.createTime | parseTime}}</span>
                </template>
            </el-table-column>
            <el-table-column fixed="right" label="操作" align="center" width="200" class-name="small-padding fixed-width">
                <template slot-scope="scope">
                <el-button v-if="scope.row.mailState != 1 || scope.row.deleteTime > 0" type="primary" size="mini" @click="handleUpdate(scope.row)">查看</el-button>
                <el-button v-if="scope.row.mailState == 1 && scope.row.deleteTime == 0" type="primary" size="mini" @click="handleUpdate(scope.row)">编辑</el-button>
                <el-button v-if="scope.row.mailState == 1 && scope.row.deleteTime == 0" size="mini" type="danger" @click="handleDelete(scope.row)">删除</el-button>
                </template>
            </el-table-column>
        </el-table>

        <div class="pagination-container" style="margin-top:15px;">
            <el-pagination :current-page="listQuery.pageIndex" :page-sizes="[20]" :total="total" background layout="total, sizes, prev, pager, next, jumper"  @current-change="handleCurrentChange"/>
        </div>

        <el-dialog :title="textMap[dialogStatus]" :visible.sync="dialogFormVisible">
            <el-form ref="dataForm" :model="temp" label-position="left" label-width="100px" style="width: 600px; margin-left:50px;">
                <el-form-item label="邮件类型">
                    <el-radio-group v-model="temp.mailType">
                        <el-radio :label="1">个人邮件</el-radio>
                        <el-radio :label="2">全服邮件</el-radio>
                    </el-radio-group>
                </el-form-item>
                <el-form-item label="标题">
                    <el-input v-model="temp.title"/>
                </el-form-item>
                <el-form-item v-if="dialogStatus=='create'" label="服务器">
                    <el-select v-model="temp.channelId" placeholder="渠道" style="width: 120px" class="filter-item" @change="handleChannelChange">
                        <el-option v-for="item in channelList" :key="item.channelId" :label="item.channelName" :value="item.channelId"/>
                    </el-select>

                    <el-select v-model="temp.platformId" placeholder="平台" clearable style="width: 160px" class="filter-item" @change="handlePlatformChange">
                        <el-option v-for="item in tempPlatformList" :key="item.platformId" :label="item.platformName" :value="item.platformId" />
                    </el-select>

                    <el-select v-model="temp.serverId" collapse-tags placeholder="服务器" clearable style="width: 180px" class="filter-item" >
                        <el-option v-for="item in serverList" :key="item.id" :label="item.serverName" :value="item.id"/>
                    </el-select>
                </el-form-item>
                <el-form-item v-if="dialogStatus=='update'" label="服务器">
                    <el-input v-model="temp.serverName" :disabled="true"/>
                </el-form-item>
                <el-form-item label="内容">
                    <el-input v-model="temp.content"/>
                </el-form-item>
                <el-form-item label="玩家列表">
                    <el-input v-model="temp.playerlist"/>
                </el-form-item>
                <el-form-item label="道具列表">
                    <el-input v-model="temp.proplist"/>
                </el-form-item>
                <el-form-item label="是否绑定">
                    <el-select v-model="temp.bindFlag" placeholder="是否绑定" style="width: 120px" class="filter-item">
                        <el-option v-for="item in yesOrNoArray" :key="item.key" :label="item.name" :value="item.key"/>
                    </el-select>
                </el-form-item>
                <!-- <el-form-item label="冻结时间(分钟)">
                    <el-input v-model="temp.freezTime" type="number"/>
                </el-form-item>
                <el-form-item label="邮件有效天数">
                    <el-input v-model="temp.effectDays" type="number"/>
                </el-form-item> -->
                <el-form-item label="角色创建开始时间">
                    <el-date-picker v-model="temp.roleStartTimeStr" type="datetime" placeholder="角色创建开始时间">
                    </el-date-picker>
                </el-form-item>
                <!-- <el-form-item label="角色创建结束时间">
                    <el-date-picker v-model="temp.roleEndTimeStr" type="datetime" placeholder="角色创建结束时间">
                    </el-date-picker>
                </el-form-item> -->
                <el-form-item label="最小等级">
                    <el-input v-model="temp.minLevel" type="number"/>
                </el-form-item>
                <el-form-item label="备注">
                    <el-input v-model="temp.remark"/>
                </el-form-item>
                <!-- <el-form-item label="最大等级">
                    <el-input v-model="temp.maxLevel" type="number"/>
                </el-form-item> -->
            </el-form>
            <div slot="footer" class="dialog-footer">
                <el-button @click="dialogFormVisible = false">取消</el-button>
                <el-button v-if="dialogStatus=='create'" type="primary" @click="createData">创建</el-button>
                <el-button v-else-if="temp.mailState == 1 && temp.deleteTime == 0" type="primary" @click="updateData">确定</el-button>
            </div>
        </el-dialog>
        <el-dialog :visible.sync="dialogPvVisible" title="是否确认删除" width="30%">
      <div>
        是否确认删除邮件：{{ temp.title }}
      </div>
      <span slot="footer" class="dialog-footer">
        <el-button @click="dialogPvVisible = false">取消</el-button>
        <el-button type="danger" @click="deleteData">删除</el-button>
      </span>
    </el-dialog>
    </div>
</template>

<script>
import waves from "@/directive/waves"; // 水波纹指令
import { parseTime } from "@/utils/index";
import { getAllChannel } from "@/api/channel";
import { getAllPlatformList } from "@/api/platform";
import { getCenterServerList } from "@/api/center";
import { Message, MessageBox } from "element-ui";
import { mailState, mailType } from "@/types/manage";
import { getApplyList, addmail, updatemail, deletemail } from "@/api/mail";
import { yesOrNoList } from "@/types/public";
import { checkItemContent } from "@/utils/tool.js";

export default {
  name: "ApplymailList",
  directives: {
    waves
  },
  filters: {
    parseTime: function(value) {
      return parseTime(value, "{y}-{m}-{d} {h}:{i}");
    },
    parseMailState: function(value) {
      if (mailState[value - 1]) {
        return mailState[value - 1].name;
      }
      return "";
    },
    parseYesOrNo: function(value) {
      if (value == 1) {
        return "是";
      }
      return "否";
    },
    parseMailType: function(value) {
      if (mailType[value - 1]) {
        return mailType[value - 1].name;
      }
      return "";
    },
    parseSecond: function(value) {
      let hour = parseInt(value / 60 / 60 / 1000);
      let reseMinute = value % (60 * 60);
      let minute = parseInt(reseMinute / 60);
      let reseSecond = reseMinute % 60;
      return hour + "时" + minute + "分" + reseSecond + "秒";
    }
  },
  created() {
    this.initMetaData();
    // this.getList();
  },
  data() {
    return {
      listLoading: false,
      mailStateList: [],
      mailTypeList: [],
      tableKey: 1,
      total: 0,
      listQuery: {
        pageIndex: 1,
        mailState: undefined,
        serverId: undefined,
        title: ""
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

      channelList: [],
      platformList: [],
      groupList: [],
      tempPlatformList: [],
      serverList: [],
      chatForbidTimeArray: [],
      monitorTemp: {},
      yesOrNoArray: []
    };
  },
  methods: {
    handleFilter: function() {
      console.log(this.listQuery);
      this.listQuery.pageIndex = 1;
      this.listQuery.ordercol = 1;
      this.listQuery.ordertype = 0;
      this.getList();
    },

    getList() {
      this.listLoading = true;
      getApplyList(this.listQuery)
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
    handleCreate() {
      this.dialogStatus = "create";
      this.dialogFormVisible = true;
      this.temp = {
        id: 0,
        mailType: 1,
        serverId: undefined,
        serverName: undefined,
        title: undefined,
        content: undefined,
        playerlist: undefined,
        proplist: undefined,
        freezTime: undefined,
        effectDays: undefined,
        roleStartTime: undefined,
        roleStartTimeStr: new Date(),
        roleEndTime: undefined,
        roleEndTimeStr: new Date(),
        minLevel: undefined,
        maxLevel: undefined,
        sdkType: undefined,
        centerPlatformId: undefined,
        bindFlag: undefined,
        remark: undefined
      };
    },
    createData() {
      let flag = this.checkCommit();
      if (!flag) {
        return;
      }
      this.temp.roleStartTime = this.temp.roleStartTimeStr.valueOf();
      this.temp.roleEndTime = this.temp.roleEndTimeStr.valueOf();
      this.temp.serverId = parseInt(this.temp.serverId);
      this.temp.mailType = parseInt(this.temp.mailType);
      this.temp.freezTime = parseInt(this.temp.freezTime);
      this.temp.effectDays = parseInt(this.temp.effectDays);
      this.temp.minLevel = parseInt(this.temp.minLevel);
      this.temp.maxLevel = parseInt(this.temp.maxLevel);
      this.temp.sdkType = parseInt(this.temp.sdkType);
      this.temp.centerPlatformId = parseInt(this.temp.centerPlatformId);
      this.temp.channelList = parseInt(this.temp.channelId);
      this.temp.platformId = parseInt(this.temp.platformId);
      this.temp.bindFlag = parseInt(this.temp.bindFlag);

      addmail(this.temp).then(() => {
        this.getList();
        this.dialogFormVisible = false;
        this.showSuccess();
      });
    },
    handleUpdate(e) {
      const curRow = Object.assign({}, e);
      //   this.temp.startTimestr = new Date(this.temp.startTime);
      this.temp = {
        id: curRow.id,
        mailType: curRow.mailType,
        serverId: curRow.serverId,
        title: curRow.title,
        content: curRow.content,
        playerlist: curRow.playerlist,
        proplist: curRow.proplist,
        freezTime: curRow.freezTime,
        effectDays: curRow.effectDays,
        roleStartTime: curRow.roleStartTime,
        serverName: curRow.serverName,
        roleStartTimeStr: new Date(curRow.roleStartTime),
        roleEndTime: curRow.roleEndTime,
        roleEndTimeStr: new Date(curRow.roleEndTime),
        minLevel: curRow.minLevel,
        maxLevel: curRow.maxLevel,
        mailState: curRow.mailState,
        sdkType: curRow.sdkType,
        centerPlatformId: curRow.centerPlatformId,
        bindFlag: curRow.bindFlag,
        deleteTime: curRow.deleteTime,
        remark: curRow.remark
      };
      this.dialogStatus = "update";
      this.dialogFormVisible = true;
    },
    updateData() {
      let flag = this.checkCommitEdit();
      if (!flag) {
        return;
      }
      this.temp.roleStartTime = this.temp.roleStartTimeStr.valueOf();
      this.temp.roleEndTime = this.temp.roleEndTimeStr.valueOf();
      this.temp.serverId = parseInt(this.temp.serverId);
      this.temp.mailType = parseInt(this.temp.mailType);
      this.temp.freezTime = parseInt(this.temp.freezTime);
      this.temp.effectDays = parseInt(this.temp.effectDays);
      this.temp.minLevel = parseInt(this.temp.minLevel);
      this.temp.maxLevel = parseInt(this.temp.maxLevel);
      this.temp.sdkType = parseInt(this.temp.sdkType);
      this.temp.centerPlatformId = parseInt(this.temp.centerPlatformId);

      updatemail(this.temp).then(() => {
        this.getList();
        this.dialogFormVisible = false;
        this.showSuccess();
      });
    },
    handleDelete(e) {
      this.dialogPvVisible = true;
      this.temp = Object.assign({}, e);
    },
    deleteData() {
      deletemail(this.temp).then(() => {
        this.getList();
        this.dialogPvVisible = false;
        this.showSuccess();
      });
    },
    initMetaData() {
      this.mailStateList = mailState;
      this.mailTypeList = mailType;
      this.yesOrNoArray = yesOrNoList;
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
        this.temp.serverId = undefined;
      }
    },
    handlePlatformChange: function(e) {
      console.log(e);
      if (e) {
        let item = this.findPlatFormItem(e);
        if (item) {
          this.temp.sdkType = item.sdkType;
          this.temp.centerPlatformId = item.centerPlatformId;
          getCenterServerList(item.centerPlatformId).then(res => {
            this.serverList = res.itemArray;
            this.temp.serverId = undefined;
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
    showSuccess() {
      this.$message({
        message: "设置成功",
        type: "success",
        duration: 1000
      });
    },
    checkCommit() {
      if (this.temp.mailType == 1) {
        if (
          !this.temp.channelId ||
          !this.temp.title ||
          !this.temp.content ||
          !this.temp.playerlist ||
          !this.temp.mailType
        ) {
          this.showError("参数不完整");
          return false;
        }
        if(!checkItemContent(this.temp.proplist)){
            this.showError("道具参数不对");
          return false;
        }
      }
      if (this.temp.mailType == 2) {
        console.log(this.temp);
        //quanfu mail
        if (
          !this.temp.channelId ||
          !this.temp.title ||
          !this.temp.content ||
          !this.temp.roleStartTimeStr ||
          !this.temp.minLevel ||
          !this.temp.mailType
        ) {
          this.showError("参数不完整");
          return false;
        }
        if(!checkItemContent(this.temp.proplist)){
            this.showError("道具参数不对");
          return false;
        }
      }
      return true;
    },
    checkCommitEdit() {
      if (this.temp.mailType == 1) {
        if (
          !this.temp.serverId ||
          !this.temp.title ||
          !this.temp.content ||
          !this.temp.playerlist ||
          !this.temp.mailType
        ) {
          this.showError("参数不完整");
          return false;
        }
        if(!checkItemContent(this.temp.proplist)){
            this.showError("道具参数不对");
          return false;
        }
      }
      if (this.temp.mailType == 2) {
        console.log(this.temp);
        //quanfu mail
        if (
          !this.temp.serverId ||
          !this.temp.title ||
          !this.temp.content ||
          !this.temp.roleStartTimeStr ||
          !this.temp.minLevel ||
          !this.temp.mailType
        ) {
          this.showError("参数不完整");
          return false;
        }
        if(!checkItemContent(this.temp.proplist)){
            this.showError("道具参数不对");
          return false;
        }
      }
      return true;
    },
    showError(msg) {
      Message({
        message: msg,
        type: "error",
        duration: 1.5 * 1000
      });
    }
  }
};
</script>

