<template>
    <div class="app-container">
        <div class="filter-container">
            <el-input placeholder="订单号" v-model="listQuery.orderId" style="width: 200px;" class="filter-item"/>
            <el-input placeholder="Sdk订单号" v-model="listQuery.sdkOrderId" style="width: 200px;" class="filter-item" @keyup.enter.native="handleFilter"/>
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
            <el-table-column label="ID" align="center" width="120px">
                <template slot-scope="scope">
                    <span>{{ scope.row.id }}</span>
                </template>
            </el-table-column>
            <el-table-column label="订单号" min-width="150px" align="left">
                <template slot-scope="scope">
                    <span>{{ scope.row.orderId}}</span>
                </template>
            </el-table-column>
            <el-table-column label="Sdk订单号" min-width="150px" align="left">
                <template slot-scope="scope">
                    <span>{{ scope.row.sdkOrderId}}</span>
                </template>
            </el-table-column>
            <el-table-column label="订单状态" min-width="150px" align="left">
                <template slot-scope="scope">
                    <span>{{ scope.row.status | parseOrderStatus}}</span>
                </template>
            </el-table-column>
            <el-table-column label="Sdk类型" min-width="150px" align="left">
                <template slot-scope="scope">
                    <span>{{ getSdkTypeName(scope.row.sdkType)}}</span>
                </template>
            </el-table-column>
            <el-table-column label="服务器序号" width="100px" align="left">
                <template slot-scope="scope">
                    <span>{{ scope.row.serverId}}</span>
                </template>
            </el-table-column>
            <el-table-column label="账户Id" min-width="150px" align="left">
                <template slot-scope="scope">
                    <span>{{ scope.row.userId}}</span>
                </template>
            </el-table-column>
            <el-table-column label="角色Id" min-width="150px" align="left">
                <template slot-scope="scope">
                    <span>{{ scope.row.playerId}}</span>
                </template>
            </el-table-column>
            <el-table-column label="角色名称" min-width="150px" align="left">
                <template slot-scope="scope">
                    <span>{{ scope.row.playerName}}</span>
                </template>
            </el-table-column>
            <el-table-column label="充值档次" min-width="100px" align="left">
                <template slot-scope="scope">
                    <span>{{ scope.row.chargeId}}</span>
                </template>
            </el-table-column>
            <el-table-column label="充值金额" min-width="100px" align="left">
                <template slot-scope="scope">
                    <span>{{ scope.row.money}}</span>
                </template>
            </el-table-column>
            <el-table-column label="接收时间" min-width="150px" align="left">
                <template slot-scope="scope">
                    <span>{{ scope.row.receivePayTime | parseTime}}</span>
                </template>
            </el-table-column>
            <el-table-column label="创建时间" min-width="150px" align="left">
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
import { getCenterOrderList } from "@/api/centerorder";
import { sdkTypeList,orderStateList } from "@/types/center";
import { parseTime } from "@/utils/index";
import { getAllSdkType } from "@/api/center";
export default {
  name: "CenterOrderList",
  directives: {
    waves
  },
  filters: {
    parseSdkType: function(value) {
      if (sdkTypeList[value - 1]) {
        return sdkTypeList[value - 1].name;
      }
      return "";
    },
    parseTime: function(value) {
      return parseTime(value, "{y}-{m}-{d} {h}:{i}:{s}");
    },
    parseOrderStatus:function(value){
        if(orderStateList[value]){
            return orderStateList[value].name
        }
        return ""
    }
  },
  created() {
    this.skdTypeArray = sdkTypeList;
    // this.getList();
    getAllSdkType().then(res =>{
        this.skdTypeArray = res.itemArray
    })
  },
  data() {
    return {
      listLoading: false,
      tableKey: 0,
      total: 0,
      listQuery: {
        pageIndex: 1,
        sdkOrderId: undefined,
        orderId: undefined
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
        sdkType: undefined,
        centerPlatformName: undefined
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
    getList() {
      getCenterOrderList(this.listQuery)
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

