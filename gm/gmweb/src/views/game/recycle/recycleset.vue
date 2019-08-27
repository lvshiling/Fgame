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
            <el-table-column label="ID" align="center" width="180px" >
                <template slot-scope="scope">
                    <span>{{ scope.row.id }}</span>
                </template>
            </el-table-column>
            <el-table-column label="服务ID" min-width="80px" align="left">
                <template slot-scope="scope">
                    <span>{{ scope.row.serverId}}</span>
                </template>
            </el-table-column>
            <el-table-column label="已回收金额" min-width="90px" align="left">
                <template slot-scope="scope">
                    <span>{{ scope.row.recycleGold}}</span>
                </template>
            </el-table-column>
            <el-table-column label="回收时间" min-width="90px" align="left">
                <template slot-scope="scope">
                    <span>{{ scope.row.recycleTime | parseTime}}</span>
                </template>
            </el-table-column>
            <el-table-column label="自定义回收金额" min-width="90px" align="left">
                <template slot-scope="scope">
                    <span>{{ scope.row.customRecycleGold}}</span>
                </template>
            </el-table-column>
            <el-table-column fixed="right" label="操作" align="center" width="200" class-name="small-padding fixed-width">
                <template slot-scope="scope">
                <el-button type="primary" size="mini" @click="handleRecyle(scope.row)">设置</el-button>
                </template>
            </el-table-column>
        </el-table>

        <el-dialog :visible.sync="dialogPvVisible" title="自定义回收金额">
          <el-form ref="dataForm" :model="monitorTemp" label-position="left" label-width="120px" style="width: 400px; margin-left:50px;">
            <el-form-item label="金额">
              <el-input-number v-model="monitorTemp.gold" label="自定义回收金额"></el-input-number>
            </el-form-item>
          </el-form>
          <div slot="footer" class="dialog-footer">
            <el-button @click="dialogPvVisible = false">取消</el-button>
            <el-button type="primary" @click="updateRecyle">设置</el-button>
          </div>
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
import { getRecycleList, setRecycleGold } from "@/api/recycle";

export default {
  name: "recycleSet",
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
        platformId: undefined,
        channelId: undefined,
        serverId: undefined
      },
      channelList: [],
      platformList: [],
      groupList: [],
      serverList: [],
      tempPlatformList: [],
      textMap: {
        update: "编辑",
        create: "添加"
      },
      dialogStatus: "",
      dialogPvVisible: false,
      monitorTemp: {
        serverId: undefined,
        gold: undefined
      },
      temp: {},
      list: []
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
      this.getList();
    },
    handleRecyle(e) {
      this.dialogPvVisible = true;
      // this.monitorTemp = e;
      this.monitorTemp.serverId = this.listQuery.serverId;
      this.monitorTemp.gold = e.customRecycleGold;
    },
    updateRecyle() {
      console.log(this.monitorTemp);
      setRecycleGold(this.monitorTemp).then(res => {
        this.showSuccess();
        this.dialogPvVisible = false;
        this.getList();
      });
    },
    getList() {
      this.listLoading = true;
      getRecycleList(this.listQuery)
        .then(res => {
          this.list = res.itemArray;
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

