<template>
    <div class="app-container">
        <div class="filter-container">
            <el-input placeholder="中心平台名" v-model="listQuery.centerPlatformName" style="width: 200px;" class="filter-item" @keyup.enter.native="handleFilter"/>
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
            <el-table-column label="中心平台ID" align="center" width="120px">
                <template slot-scope="scope">
                    <span>{{ scope.row.centerPlatformId }}</span>
                </template>
            </el-table-column>
            <el-table-column label="中心平台名" min-width="150px" align="left">
                <template slot-scope="scope">
                    <span>{{ scope.row.centerPlatformName}}</span>
                </template>
            </el-table-column>
            <!-- <el-table-column label="Sdk类型" min-width="150px" align="left">
                <template slot-scope="scope">
                    <span>{{ scope.row.sdkType | parseSdkType}}</span>
                </template>
            </el-table-column> -->
            <el-table-column label="操作" align="center" width="280" class-name="small-padding fixed-width">
                <template slot-scope="scope">
                <el-button type="primary" size="mini" @click="handleUpdate(scope.row)">编辑</el-button>
                <el-button size="mini" type="danger" @click="handleDelete(scope.row)">删除</el-button>
                <el-button type="primary" size="mini" @click="handleRefresh(scope.row)">刷新</el-button>
                </template>
            </el-table-column>
        </el-table>

        <div class="pagination-container" style="margin-top:15px;">
            <el-pagination :current-page="listQuery.pageIndex" :page-sizes="[20]" :total="total" background layout="total, sizes, prev, pager, next, jumper"  @current-change="handleCurrentChange"/>
        </div>

        <el-dialog :title="textMap[dialogStatus]" :visible.sync="dialogFormVisible">
        <el-form ref="dataForm" :model="temp" label-position="left" label-width="120px" style="width: 400px; margin-left:50px;">
            <el-form-item label="中心平台名">
                <el-input v-model="temp.centerPlatformName"/>
            </el-form-item>
            <!-- <el-form-item label="sdk类型">
              <el-select v-model="temp.sdkType" placeholder="sdk类型" style="width: 160px" class="filter-item" >
                <el-option v-for="item in skdTypeArray" :key="item.key" :label="item.name" :value="item.key" />
              </el-select>
            </el-form-item> -->
        </el-form>
        <div slot="footer" class="dialog-footer">
            <el-button @click="dialogFormVisible = false">取消</el-button>
            <el-button v-if="dialogStatus=='create'" type="primary" @click="createData">创建</el-button>
            <el-button v-else type="primary" @click="updateData">确定</el-button>
        </div>
        </el-dialog>

        <el-dialog :visible.sync="dialogPvVisible" title="是否确认删除" width="30%">
          <div>
              是否确认删除中心平台：{{temp.centerPlatformName}}
          </div>
          <span slot="footer" class="dialog-footer">
            <el-button @click="dialogPvVisible = false">取消</el-button>
            <el-button type="danger" @click="deleteData">删除</el-button>
          </span>
        </el-dialog>

        <el-dialog :visible.sync="dialogRefreshVisible" title="是否确认刷新" width="30%">
          <div>
            是否确认刷新中心服务：{{ temp.serverName }}
          </div>
          <span slot="footer" class="dialog-footer">
            <el-button @click="dialogRefreshVisible = false">取消</el-button>
            <el-button type="danger" @click="refreshData">刷新</el-button>
          </span>
        </el-dialog>
        
    </div>
</template>

<script>
import waves from "@/directive/waves"; // 水波纹指令
import {
  getCenterPlatformList,
  getCenterPlatformAdd,
  getCenterPlatformUpdate,
  getCenterPlatformDelete
} from "@/api/centerPlatform";
import { refreshCenterServer } from "@/api/centerserver";
import { sdkTypeList } from "@/types/center";
import { getAllSdkType } from "@/api/center";
export default {
  name: "centerPlatformList",
  directives: {
    waves
  },
  filters: {
    parseSdkType: function(value) {
      if(sdkTypeList[value - 1]){
        return sdkTypeList[value - 1].name;
      }
      return ''
    }
  },
  created() {
    this.skdTypeArray = sdkTypeList;
    this.getList();
    getAllSdkType().then(res =>{
        this.skdTypeArray = res.itemArray
    });
  },
  data() {
    return {
      listLoading: false,
      tableKey: 0,
      total: 0,
      listQuery: {
        pageIndex: 1,
        centerPlatformName: ""
      },
      textMap: {
        update: "编辑",
        create: "添加"
      },
      dialogStatus: "",
      dialogPvVisible: false,
      dialogFormVisible: false,
      dialogRefreshVisible: false,
      temp: {
        sdkType:undefined,
        centerPlatformName:undefined
      },
      list: [],
      skdTypeArray: []
    };
  },
  methods: {
    handleFilter: function() {
      this.listQuery.pageIndex = 1;
      this.getList();
    },
    handleCreate: function() {
      this.dialogStatus = "create";
      this.dialogFormVisible = true;
      this.temp = {
        sdkType:undefined,
        centerPlatformName:undefined
      };
    },
    handleUpdate: function(e) {
      this.dialogStatus = "update";
      this.dialogFormVisible = true;
      this.temp = Object.assign({}, e);
    },

    handleDelete: function(e) {
      this.dialogPvVisible = true;
      this.temp = Object.assign({}, e);
    },
    getList() {
      this.listLoading = true;
      let centerPlatformName = this.listQuery.centerPlatformName;
      let pageIndex = this.listQuery.pageIndex;
      getCenterPlatformList(centerPlatformName, pageIndex)
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
      this.listQuery.pageIndex = e;
      this.getList();
    },

    updateData() {
      getCenterPlatformUpdate(this.temp).then(() => {
        this.getList();
        this.dialogFormVisible = false;
        this.showSuccess();
      });
    },
    createData() {
      getCenterPlatformAdd(this.temp).then(() => {
        this.getList();
        this.dialogFormVisible = false;
        this.showSuccess();
      });
    },
    deleteData() {
      getCenterPlatformDelete(this.temp).then(() => {
        this.getList();
        this.dialogPvVisible = false;
        this.showSuccess();
      });
    },
    handleRefresh: function(e) {
      const curRow = Object.assign({}, e);
      this.temp = curRow;
      this.dialogStatus = "refresh";
      this.dialogRefreshVisible = true;
    },
    refreshData() {
      let myserverId = parseInt(this.temp.centerPlatformId);
      let postData = {
        serverId: myserverId
      };
      refreshCenterServer(postData).then(() => {
        this.dialogRefreshVisible = false;
        this.showSuccess();
      });
    },
    showSuccess() {
      this.$message({
        message: "修改成功",
        type: "success",
        duration: 1000
      });
    },
    getSdkTypeName(value){
        for(let i=0,len=this.skdTypeArray.length;i<len;i++){
            if(this.skdTypeArray[i].key == value){
                return this.skdTypeArray[i].name
            }
        }
        return value
    }
  }
};
</script>

