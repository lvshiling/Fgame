<template>
    <div class="app-container">
        <div>
            <div class="filter-container">
                <el-select v-model="listQuery.platformId" placeholder="平台" clearable style="width: 160px" class="filter-item" @change="handlePlatformChange">
                    <el-option v-for="item in platformList" :key="item.platformId" :label="item.platformName" :value="item.platformId" />
                </el-select>
                <el-button v-waves class="filter-item" type="primary" icon="el-icon-search" @click="handleFilter">搜索</el-button>
                <el-button class="filter-item" style="margin-left: 10px;" type="primary" icon="el-icon-edit" @click="handleCreate">添加</el-button>
            </div>
        </div>

        <el-table
            :key="chatSetKey"
            :data="chatSetList"
            v-loading="listLoading"
            border
            fit
            highlight-current-row
            style="width: 100%;margin-top:15px;">
            <el-table-column label="配置Id" align="center" width="180px">
                <template slot-scope="scope">
                    <span>{{ scope.row.chatSetId }}</span>
                </template>
            </el-table-column>
            <el-table-column label="平台Id" align="center" width="100px">
                <template slot-scope="scope">
                    <span>{{ scope.row.centerPlatformId }}</span>
                </template>
            </el-table-column>
            <el-table-column label="平台名" align="center" width="180px">
                <template slot-scope="scope">
                    <span>{{ findPlatFormName(scope.row.centerPlatformId) }}</span>
                </template>
            </el-table-column>
            <el-table-column label="世界VIP等级" width="100px" align="center">
                <template slot-scope="scope">
                    <span>{{ scope.row.worldVip}}</span>
                </template>
            </el-table-column>
            <el-table-column label="世界玩家等级" width="110px">
                <template slot-scope="scope">
                <span>{{ scope.row.worldPlayerLevel}}</span>
                </template>
            </el-table-column>
            
            <el-table-column label="私聊VIP等级" width="100px" align="center">
                <template slot-scope="scope">
                    <span>{{ scope.row.pChatVip}}</span>
                </template>
            </el-table-column>
            <el-table-column label="私聊玩家等级" width="110px">
                <template slot-scope="scope">
                <span>{{ scope.row.pChatPlayerLevel}}</span>
                </template>
            </el-table-column>
            
            <el-table-column label="公会VIP等级" width="100px" align="center">
                <template slot-scope="scope">
                    <span>{{ scope.row.guildVip}}</span>
                </template>
            </el-table-column>
            <el-table-column label="公会玩家等级" width="110px">
                <template slot-scope="scope">
                <span>{{ scope.row.guildPlayerLevel}}</span>
                </template>
            </el-table-column>
            
            <el-table-column label="组队vip等级" width="100px" align="center">
                <template slot-scope="scope">
                    <span>{{ scope.row.teamVip}}</span>
                </template>
            </el-table-column>
            <el-table-column label="组队玩家等级" width="110px">
                <template slot-scope="scope">
                <span>{{ scope.row.teamPlayerLevel}}</span>
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
            <el-form ref="dataForm" :model="temp" label-position="left" label-width="100px" style="width: 400px; margin-left:50px;">
                <el-form-item label="平台">
                    <el-select v-model="temp.platformId" placeholder="平台" :disabled="dialogStatus=='update'" style="width: 160px" class="filter-item"  @change="handleItemPlatformChange">
                        <el-option v-for="item in platformList" :key="item.platformId" :label="item.platformName" :value="item.platformId" />
                    </el-select>
                </el-form-item>

                <el-form-item label="世界VIP等级">
                    <el-input type="number" v-model="temp.worldVip"/>
                </el-form-item>

                <el-form-item label="世界玩家等级">
                    <el-input type="number" v-model="temp.worldPlayerLevel"/>
                </el-form-item>
                
                <el-form-item label="私聊VIP等级">
                    <el-input type="number" v-model="temp.pChatVip"/>
                </el-form-item>

                <el-form-item label="私聊玩家等级">
                    <el-input type="number" v-model="temp.pChatPlayerLevel"/>
                </el-form-item>
                
                <el-form-item label="公会VIP等级">
                    <el-input type="number" v-model="temp.guildVip"/>
                </el-form-item>

                <el-form-item label="工会玩家等级">
                    <el-input type="number" v-model="temp.guildPlayerLevel"/>
                </el-form-item>

                <el-form-item label="组队vip等级">
                    <el-input type="number" v-model="temp.teamVip"/>
                </el-form-item>

                <el-form-item label="组队玩家等级">
                    <el-input type="number" v-model="temp.teamPlayerLevel"/>
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
import {
  addChatSetPlatform,
  updateChatSetPlatform,
  deleteChatSetPlatform,
  getChatSetListPlatform
} from "@/api/chatset";
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
      chatSetList: [],
      dialogFormVisible: false,
      dialogPvVisible: false,
      dialogStatus: undefined,
      textMap: {
        update: "编辑",
        create: "添加"
      },
      listQuery: {
        centerPlatformId: undefined,
        centerServerId: undefined,
        pageIndex: 0
      },
      platformList: [],
      serverList: [],
      temp: {
        startTime: undefined,
        endTime: undefined,
        centerPlatformId: undefined,
        centerServerId: undefined,
        centerServerArray:[]
      },
      tempServerList: []
    };
  },
  methods: {
    initMetaData() {
      getAllPlatformList().then(res => {
        this.platformList = res.itemArray;
      });
      // this.loadData();
    },
    handleFilter(e) {
      if(!this.listQuery.centerPlatformId){
        this.showChose("请选择平台")
        return
      }
      this.listQuery.pageIndex = 1;
      this.loadData();
    },
    handleCreate(e) {
      this.temp = {
        startTime: undefined,
        endTime: undefined,
        centerPlatformId: undefined,
        centerServerId: undefined,
        centerServerArray: [],
        skdType : undefined
      };
      this.dialogFormVisible = true;
      this.dialogStatus = "create";
      this.tempServerList = [];
    },
    handleEdit(e) {
      this.dialogFormVisible = true;
      this.dialogStatus = "update";
      this.temp = Object.assign({}, e);
      this.temp.startTime = new Date("2000-01-01 " + this.temp.startTime);
      this.temp.endTime = new Date("2000-01-01 " + this.temp.endTime);
      let item = this.findPlatFormItemByCenterId(this.temp.centerPlatformId)
      if(item){
        this.temp.platformId = item.platformId
      }
      getCenterServerList(this.temp.centerPlatformId).then(res => {
        this.tempServerList = res.itemArray;
      });
      this.temp.centerServerArray = [this.temp.centerServerId]
    },
    handleDelete(e) {
      this.dialogPvVisible = true;
      this.temp = Object.assign({}, e);
    },
    handleCurrentChange(e) {
      this.listQuery.pageIndex = e;
      this.loadData();
    },
    handlePlatformChange(e) {
      if (!e || e == "" || e == undefined) {
        this.listQuery.centerPlatformId = undefined;
        this.listQuery.centerServerId = undefined;
        return;
      }
      this.listQuery.centerServerId = undefined;
      let platformItem = this.findPlatFormItem(this.listQuery.platformId);
      if (platformItem) {
        this.listQuery.centerPlatformId = platformItem.centerPlatformId;
        // getCenterServerList(platformItem.centerPlatformId).then(res => {
        //   this.serverList = res.itemArray;
        // });
      }
    },
    handleItemPlatformChange(e) {
      this.temp.centerServerId = undefined;
      let platformItem = this.findPlatFormItem(this.temp.platformId);
      if (platformItem) {
        this.temp.centerPlatformId = platformItem.centerPlatformId;
        this.temp.sdkType = platformItem.sdkType;
        getCenterServerList(this.temp.centerPlatformId).then(res => {
          this.tempServerList = res.itemArray;
        });
      }
    },
    findPlatFormName(platformId) {
      let platform = this.platformList.find(n => {
        return n.centerPlatformId == platformId;
      });
      if (platform) {
        return platform.platformName;
      }
      return undefined;
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
    findPlatFormItemByCenterId(centerPlatformId) {
      let platform = this.platformList.find(n => {
        return n.centerPlatformId == centerPlatformId;
      });
      if (platform) {
        return platform;
      }
      return undefined;
    },
    createData(e) {
      addChatSetPlatform(this.temp).then(res => {
        this.dialogFormVisible = false;
        this.showSuccess();
        this.loadData();
      });
    },
    updateData(e) {
      updateChatSetPlatform(this.temp).then(res => {
        this.dialogFormVisible = false;
        this.showSuccess();
        this.loadData();
      });
    },
    deleteData(e) {
      deleteChatSetPlatform(this.temp).then(res => {
        this.dialogPvVisible = false;
        this.showSuccess();
        this.loadData();
      });
    },
    loadData() {
      this.listLoading = true;
      getChatSetListPlatform(this.listQuery).then(res => {
        this.chatSetList = res.itemArray;
        this.total = res.total;
        this.listLoading = false;
      });
    },
    showSuccess() {
      this.$message({
        message: "修改成功",
        type: "success",
        duration: 1000
      });
    },
    showChose(msg){
      this.$message({
        message: msg,
        type: "error",
        duration: 1000
      });
    }
  }
};
</script>

