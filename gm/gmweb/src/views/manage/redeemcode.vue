<template>
    <div class="app-container">
        <div class="filter-container">
            <el-input placeholder="礼包id" v-model="listQuery.id" style="width: 200px;" class="filter-item"/>
            <el-button v-waves class="filter-item" type="primary" icon="el-icon-search" @click="handleFilter">搜索</el-button>
            <el-button v-waves class="filter-item" type="primary" icon="el-icon-search" @click="handleDown">下载</el-button>
        </div>

        <el-table
            v-loading="listLoading"
            :key="tableKey"
            :data="list"
            border
            fit
            highlight-current-row
            style="width: 100%;margin-top:15px;">
            <el-table-column
            type="index"
            width="100">
            </el-table-column>
            <!-- <el-table-column label="ID" align="center" width="120px">
                <template slot-scope="scope">
                    <span>{{ scope.row.id }}</span>
                </template>
            </el-table-column> -->
             <el-table-column label="兑换码" align="center" width="300px">
                <template slot-scope="scope">
                    <span>{{ scope.row.redeemCode }}</span>
                </template>
            </el-table-column>
            <el-table-column label="已使用次数" align="center" width="120px">
                <template slot-scope="scope">
                    <span>{{ scope.row.useNum }}</span>
                </template>
            </el-table-column>
        </el-table>

        <!-- <div class="pagination-container" style="margin-top:15px;">
            <el-pagination :current-page="listQuery.pageIndex" :page-sizes="[20]" :total="total" background layout="total, sizes, prev, pager, next, jumper"  @current-change="handleCurrentChange"/>
        </div> -->
    </div>
</template>

<script>
import waves from "@/directive/waves"; // 水波纹指令
import { getRedeemCodeList,getRedeemCodeListExport } from "@/api/redeem";
import { sdkTypeList,orderStateList } from "@/types/center";
import { parseTime } from "@/utils/index";
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
    const id = this.$route.params && this.$route.params.id
    this.listQuery.id = id
    if(this.listQuery.id){
        this.getList();
    }
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
    handleDown:function(){
        this.listQuery.pageIndex = 1;
      this.exportList();
    },
    getList() {
      getRedeemCodeList(this.listQuery)
        .then(res => {
          this.list = res.itemArray;
          this.total = res.total;
          this.listLoading = false;
        })
        .catch(() => {
          this.listLoading = false;
        });
    },
    exportList() {
      getRedeemCodeListExport(this.listQuery)
        .then(res => {
          this.listLoading = false;
          if (!res) {
                return
            }
            let url = window.URL.createObjectURL(new Blob([res.data]))
            let link = document.createElement('a')
            link.style.display = 'none'
            link.href = url
            link.setAttribute('download', '兑换码.xlsx')

            document.body.appendChild(link)
            link.click()
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
    showFailMsg(msg) {
      this.$message({
        message: msg,
        type: "error",
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

