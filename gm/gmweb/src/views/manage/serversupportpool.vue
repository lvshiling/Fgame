<template>
    <div class="app-container">
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
                <el-button class="filter-item" style="margin-left: 10px;" type="primary" icon="el-icon-edit" @click="handleCreate">添加</el-button>
            </div>
        </div>

        <el-table
            :key="chatSetKey"
            :data="dataList"
            v-loading="listLoading"
            border
            fit
            highlight-current-row
            style="width: 100%;margin-top:15px;">
            <el-table-column label="扶持Id" align="center" width="180px">
                <template slot-scope="scope">
                    <span>{{ scope.row.id }}</span>
                </template>
            </el-table-column>
            <el-table-column label="服务器主键Id" align="center" width="100px">
                <template slot-scope="scope">
                    <span>{{ scope.row.serverId }}</span>
                </template>
            </el-table-column>
            <el-table-column label="服务器名称" align="center" width="180px">
                <template slot-scope="scope">
                    <span>{{ scope.row.serverName }}</span>
                </template>
            </el-table-column>
            <el-table-column label="起始元宝" align="center" width="100px">
                <template slot-scope="scope">
                    <span>{{ scope.row.beginGold }}</span>
                </template>
            </el-table-column>
            <el-table-column label="当前元宝" align="center" min-width="180px">
                <template slot-scope="scope">
                    <span>{{ scope.row.curGold }}</span>
                </template>
            </el-table-column>
            <el-table-column label="消耗元宝" align="center" min-width="180px">
                <template slot-scope="scope">
                    <span>{{ scope.row.delGold }}</span>
                </template>
            </el-table-column>

            <el-table-column label="订单扶持比例(%)" align="center" min-width="200px">
                <template slot-scope="scope">
                    <span>{{ scope.row.percent }}%</span>
                </template>
            </el-table-column>
            
            <el-table-column label="操作" align="center" width="260" class-name="small-padding fixed-width">
                <template slot-scope="scope">
                <el-button type="primary" size="mini" @click="handleEdit(scope.row)">编辑</el-button>
                <el-button size="mini" type="danger" @click="handleDelete(scope.row)">删除</el-button>
                </template>
            </el-table-column>
        </el-table>
        <div class="pagination-container" style="margin-top:15px;">
            <el-pagination :current-page="listQuery.pageIndex" :page-sizes="[20]" :total="total" background layout="total, sizes, prev, pager, next, jumper"  @current-change="handleCurrentChange"/>
        </div>

        <el-dialog :title="textMap[dialogStatus]" :visible.sync="dialogFormVisible">
            <el-form ref="dataForm" :model="temp" label-position="left" label-width="120px" style="width: 600px; margin-left:50px;">
                <el-form-item v-if="dialogStatus=='create'" label="服务器">
                    <el-select v-model="temp.channelId" placeholder="渠道" style="width: 120px" class="filter-item" @change="handleEditChannelChange">
                        <el-option v-for="item in channelList" :key="item.channelId" :label="item.channelName" :value="item.channelId"/>
                    </el-select>

                    <el-select v-model="temp.platformId" placeholder="平台" style="width: 160px" class="filter-item" @change="handleEditPlatformChange">
                        <el-option v-for="item in editPlatformList" :key="item.platformId" :label="item.platformName" :value="item.platformId" />
                    </el-select>

                    <el-select v-model="temp.serverId" collapse-tags placeholder="服务器" clearable style="width: 180px" class="filter-item" >
                        <el-option v-for="item in editServerList" :key="item.id" :label="item.serverName" :value="item.id"/>
                    </el-select>
                </el-form-item>

                <el-form-item v-if="dialogStatus=='update'" label="服务器">
                    <el-input v-model="temp.serverName" :disabled="true"/>
                </el-form-item>

                <el-form-item label="当前元宝">
                    <el-input type="number" v-model="temp.curGold"/>
                </el-form-item>
                <el-form-item label="订单扶持比例(%)，填整数">
                    <el-input type="number" v-model="temp.percent"/>
                </el-form-item>
            </el-form>

            <div slot="footer" class="dialog-footer">
                <el-button @click="dialogFormVisible = false">取消</el-button>
                <el-button v-if="dialogStatus=='create'" type="primary" @click="createData">创建</el-button>
                <el-button v-else type="primary" @click="updateData">确定</el-button>
            </div>
        </el-dialog>

        <el-dialog :visible.sync="dialogPvVisible" title="是否确认删除" width="30%">
        <div>
                是否确认删除聊天配置：
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
import { getCenterPlatList } from "@/api/center";
import { getCenterServerList } from "@/api/center";
import { getAllPlatformList } from "@/api/platform";
import { getAllChannel } from "@/api/channel";
import {
  addServerSupportPool,
  getServerSupportPoolList,
  updateServerSupportPool,
  deleteServerSupportPool
} from "@/api/supportpool";
export default {
  name: "ChatSet",
  directives: {
    waves
  },
  created() {
    this.initMetaData();
  },
  mounted() {
    //   this.websocket.close()
  },
  beforeDestroy() {},
  destroyed() {},
  data() {
    return {
      chatSetKey: 1,
      listLoading: false,
      total: 0,
      dataList: [],
      dialogFormVisible: false,
      dialogPvVisible: false,
      dialogStatus: undefined,
      textMap: {
        update: "编辑",
        create: "添加"
      },
      listQuery: {
        channelId: undefined,
        platformId: undefined,
        serverId: undefined,
        centerPlatformId: undefined,
        pageIndex: 0
      },
      serverList: [],
      temp: {
        startTime: undefined,
        endTime: undefined,
        centerPlatformId: undefined,
        centerServerId: undefined
      },
      tempServerList: [],
      channelList: [],
      platformList: [],
      tempPlatformList: [],
      editPlatformList: [],
      editServerList: []
    };
  },
  methods: {
    initMetaData() {
      getAllChannel().then(res => {
        this.channelList = res.itemArray;
      });
      getAllPlatformList().then(res => {
        this.platformList = res.itemArray;
      });
    },
    handleFilter(e) {
      this.listQuery.pageIndex = 1;
      this.loadData();
    },
    handleCreate(e) {
      this.temp = {
        serverId: undefined,
        curGold: undefined
      };
      this.dialogFormVisible = true;
      this.dialogStatus = "create";
      this.tempServerList = [];
    },
    handleEdit(e) {
      this.dialogFormVisible = true;
      this.dialogStatus = "update";
      this.temp = Object.assign({}, e);
    },
    handleDelete(e) {
      this.dialogPvVisible = true;
      this.temp = Object.assign({}, e);
    },
    handleCurrentChange(e) {
      this.listQuery.pageIndex = e;
      this.loadData();
    },
    createData(e) {
      addServerSupportPool(this.temp).then(res => {
        this.dialogFormVisible = false;
        this.showSuccess();
        this.loadData();
      });
    },
    updateData(e) {
      updateServerSupportPool(this.temp).then(res => {
        this.dialogFormVisible = false;
        this.showSuccess();
        this.loadData();
      });
    },
    deleteData(e) {
      deleteServerSupportPool(this.temp).then(res => {
        this.dialogPvVisible = false;
        this.showSuccess();
        this.loadData();
      });
    },
    loadData() {
      this.listLoading = true;
      getServerSupportPoolList(this.listQuery).then(res => {
        this.dataList = res.itemArray;
        this.total = res.total;
        this.listLoading = false;
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
          this.listQuery.centerPlatformId = item.centerPlatformId;
          getCenterServerList(item.centerPlatformId).then(res => {
            this.serverList = res.itemArray;
            this.listQuery.serverId = undefined;
          });
        }
      } else {
        this.listQuery.centerPlatformId = undefined
      }
    },
    handleEditChannelChange: function(e) {
      if (e) {
        this.temp.platformId = undefined;
        this.editPlatformList = this.findPlatFormList(e);
        if (this.editPlatformList && this.editPlatformList.length > 0) {
          // this.temp.platformId = this.editPlatformList[0].platformId;
        }
        this.groupList = [];
        this.temp.serverId = undefined;
      }
    },
    handleEditPlatformChange: function(e) {
      console.log(e);
      this.temp.serverId = undefined;
      if (e) {
        let item = this.findPlatFormItem(e);

        if (item) {
          this.temp.sdkType = item.sdkType;
          this.temp.centerPlatformId = item.centerPlatformId;
          getCenterServerList(item.centerPlatformId).then(res => {
            this.editServerList = res.itemArray;
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
        message: "修改成功",
        type: "success",
        duration: 1000
      });
    }
  }
};
</script>

