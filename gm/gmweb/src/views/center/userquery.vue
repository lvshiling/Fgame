<template>
    <div class="app-container">
        <div class="filter-container">
            <el-select v-model="listQuery.channelId" placeholder="渠道" style="width: 120px" class="filter-item" @change="handleChannelChange">
                <el-option v-for="item in channelList" :key="item.channelId" :label="item.channelName" :value="item.channelId"/>
            </el-select>

            <el-select v-model="listQuery.sdkType" placeholder="平台" clearable style="width: 160px" class="filter-item" >
                <el-option v-for="item in tempPlatformList" :key="item.sdkType" :label="item.platformName" :value="item.sdkType" />
            </el-select>

            <el-input placeholder="用户id" v-model="listQuery.userId" style="width: 200px;" class="filter-item" @keyup.enter.native="handleFilter"/>
            <el-input placeholder="平台用户Id" v-model="listQuery.platformUserId" style="width: 200px;" class="filter-item" @keyup.enter.native="handleFilter"/>
            
            <el-input placeholder="用户昵称" v-model="listQuery.userName" style="width: 200px;" class="filter-item" @keyup.enter.native="handleFilter"/>

            <el-button v-waves class="filter-item" type="primary" icon="el-icon-search" @click="handleFilter">搜索</el-button>
        </div>

        <el-table
            v-loading="listLoading"
            :key="tableKey"
            :data="list"
            border
            fit
            highlight-current-row
            style="width: 100%;margin-top:15px;">
            <el-table-column  label="用户Id" align="center" width="180px">
                <template slot-scope="scope">
                    <span>{{ scope.row.id }}</span>
                </template>
            </el-table-column>
            <el-table-column  label="中心平台Id" width="150px" align="left" >
                <template slot-scope="scope">
                    <span>{{ scope.row.platform}}</span>
                </template>
            </el-table-column>
            <el-table-column  label="平台用户Id" width="120px" align="left">
                <template slot-scope="scope">
                    <span>{{ scope.row.platformUserId}}</span>
                </template>
            </el-table-column>
            <el-table-column  label="用户昵称" width="200px" align="left">
                <template slot-scope="scope">
                    <span>{{ scope.row.name}}</span>
                </template>
            </el-table-column>
            <el-table-column  label="gm状态" width="120px" align="left">
                <template slot-scope="scope">
                    <span>{{ scope.row.gm | parseYesOrNo}}</span>
                </template>
            </el-table-column>
            <el-table-column  label="用户电话" width="120px" align="left" >
                <template slot-scope="scope">
                    <span>{{ scope.row.phoneNum}}</span>
                </template>
            </el-table-column>
            <el-table-column  label="身份证号码" min-width="150px" align="left">
                <template slot-scope="scope">
                    <span>{{ scope.row.idCard}}</span>
                </template>
            </el-table-column>
            <el-table-column  label="真实姓名" width="120px" align="left" >
                <template slot-scope="scope">
                    <span>{{ scope.row.realName}}</span>
                </template>
            </el-table-column>
            <el-table-column  label="认证状态" width="120px" align="left" >
                <template slot-scope="scope">
                    <span>{{ scope.row.realNameState | parseYesOrNo}}</span>
                </template>
            </el-table-column>
            <el-table-column  label="创建时间" width="180px" align="left">
                <template slot-scope="scope">
                    <span>{{ scope.row.createTime | parseTime}}</span>
                </template>
            </el-table-column>
        </el-table>

        <div class="pagination-container" style="margin-top:15px;">
            <el-pagination :current-page="listQuery.pageIndex" :page-sizes="[20]" :total="total" background layout="total, sizes, prev, pager, next, jumper"  @current-change="handleCurrentChange"/>
        </div>
    </div>
</template>

<script>
import waves from "@/directive/waves"; // 水波纹指令
import { parseTime } from "@/utils/index";
import { getAllChannel } from "@/api/channel";
import { getAllPlatformList } from "@/api/platform";
import { getCenterServerList } from "@/api/center";
import { Message, MessageBox } from "element-ui";
import { getCenterUserList } from "@/api/centeruser";

export default {
  name: "CenterUserQuery",
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
    parseYesOrNo:function(value){
        if(value == 1){
            return "是"
        }
        return "否"
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
        allianceName: "",
        platformUserId:"",
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
      channelList: [],
      platformList: [],
      groupList: [],
      tempPlatformList: [],
      serverList: [],
      chatForbidTimeArray: [],
      monitorTemp: {}
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
      if(this.listQuery.platformId){
          let item = this.findPlatFormItem(this.listQuery.platformId)
          if(item.centerPlatformId){
              this.listQuery.centerPlatformId = item.centerPlatformId
          }
      }else{
          this.listQuery.centerPlatformId = undefined
      }
      console.log(this.listQuery)
      this.listLoading = true;
      getCenterUserList(this.listQuery)
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
    initMetaData() {
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
            this.serverList = res.itemArray;
            this.listQuery.serverId = undefined;
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
    }
  }
};
</script>

