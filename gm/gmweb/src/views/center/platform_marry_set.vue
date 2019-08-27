<template>
    <div>
        <div class="filter-container">
            <el-input placeholder="中心平台名" v-model="listQuery.centerPlatformName" style="width: 200px;" class="filter-item" @keyup.enter.native="handleFilter"/>
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
            <el-table-column label="结婚价格数据版本" width="150px" align="left">
                <template slot-scope="scope">
                    <span>{{ scope.row.kindType | parseMarryKindType}}</span>
                </template>
            </el-table-column>
            <el-table-column label="是否设置价格" width="150px" align="left">
                <template slot-scope="scope">
                    <span>{{ scope.row.priceSetFlag | parseYesOrNo}}</span>
                </template>
            </el-table-column>
            <el-table-column label="操作" align="center" width="280px" class-name="small-padding fixed-width">
                <template slot-scope="scope">
                <el-button type="primary" size="mini" @click="handleFlag(scope.row)">设置类型</el-button>
                <!-- <el-button type="primary" size="mini" @click="handleContent(scope.row)">设置价格</el-button> -->
                </template>
            </el-table-column>
        </el-table>

        <div class="pagination-container" style="margin-top:15px;">
            <el-pagination :current-page="listQuery.pageIndex" :page-sizes="[20]" :total="total" background layout="total, sizes, prev, pager, next, jumper"  @current-change="handleCurrentChange"/>
        </div>

        <el-dialog :title="textMap[dialogStatus]" :visible.sync="dialogFormVisible">
        <el-form ref="dataForm" :model="temp" label-position="left" label-width="120px" style="width: 400px; margin-left:50px;">
            <el-form-item label="中心平台名">
                <el-input v-model="temp.centerPlatformName" disabled/>
            </el-form-item>
            <el-form-item label="价格版本类型">
              <el-select v-model="temp.kindType" placeholder="价格版本类型" style="width: 160px" class="filter-item" >
                <el-option v-for="item in kindTypeArray" :key="item.key" :label="item.name" :value="item.key" />
              </el-select>
            </el-form-item>
        </el-form>
        <div slot="footer" class="dialog-footer">
            <el-button @click="dialogFormVisible = false">取消</el-button>
            <el-button type="primary" @click="updateflag">确定</el-button>
        </div>
        </el-dialog>


        <el-dialog title="注意:此设置后廉价版和当前版本都生效" :visible.sync="dialogContentVisible">
        <el-form ref="dataForm" :model="temp" label-position="left" label-width="140px" style="width: 700px; margin-left:10px;">
            <el-form-item label="中心平台名">
                <el-input v-model="temp.centerPlatformName" disabled/>
            </el-form-item>
            <el-form-item label="当前版本婚宴价格">
                <el-row :gutter="21">
                    <el-col :span="7"><el-input v-model="marryPrice.price_1_1_1" type="number" placeholder="相濡以沫"></el-input></el-col>
                    <el-col :span="7"><el-input v-model="marryPrice.price_1_1_2" type="number" placeholder="鸾凤和鸣"></el-input></el-col>
                    <el-col :span="7"><el-input v-model="marryPrice.price_1_1_3" type="number" placeholder="龙腾凤翔"></el-input></el-col>
                </el-row>
            </el-form-item>
            <el-form-item label="当前版本戒指价格">
                <el-row :gutter="21">
                    <el-col :span="7"><el-input v-model="marryPrice.price_1_4_1" type="number" placeholder="青铜对戒"></el-input></el-col>
                    <el-col :span="7"><el-input v-model="marryPrice.price_1_4_2" type="number" placeholder="紫金对戒"></el-input></el-col>
                    <el-col :span="7"><el-input v-model="marryPrice.price_1_4_3" type="number" placeholder="龙凤对戒"></el-input></el-col>
                </el-row>
            </el-form-item>
            <el-form-item label="廉价版本婚宴价格">
                <el-row :gutter="21">
                    <el-col :span="7"><el-input v-model="marryPrice.price_2_1_1" type="number" placeholder="相濡以沫"></el-input></el-col>
                    <el-col :span="7"><el-input v-model="marryPrice.price_2_1_2" type="number" placeholder="鸾凤和鸣"></el-input></el-col>
                    <el-col :span="7"><el-input v-model="marryPrice.price_2_1_3" type="number" placeholder="龙腾凤翔"></el-input></el-col>
                </el-row>
            </el-form-item>
            <el-form-item label="廉价版本戒指价格">
                <el-row :gutter="21">
                    <el-col :span="7"><el-input v-model="marryPrice.price_2_4_1" type="number" placeholder="青铜对戒"></el-input></el-col>
                    <el-col :span="7"><el-input v-model="marryPrice.price_2_4_2" type="number" placeholder="紫金对戒"></el-input></el-col>
                    <el-col :span="7"><el-input v-model="marryPrice.price_2_4_3" type="number" placeholder="龙凤对戒"></el-input></el-col>
                </el-row>
            </el-form-item>
        </el-form>
        <div slot="footer" class="dialog-footer">
            <el-button @click="dialogContentVisible = false">取消</el-button>
            <el-button type="primary" @click="updatecontent">确定</el-button>
        </div>
        </el-dialog>
        
    </div>
</template>

<script>
import waves from "@/directive/waves"; // 水波纹指令
import {
  getCenterPlatformMarrySetList,
  updateCenterPlatformMarrySetFlag,
  updateCenterPlatformMarrySetContent
} from "@/api/centerPlatform";
import { refreshCenterServer } from "@/api/centerserver";
import { kindTypeList } from "@/types/center";
import { getAllSdkType } from "@/api/center";
import { Switch } from "element-ui";
export default {
  name: "centerPlatformMarryList",
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
    parseMarryKindType: function(value) {
      if (value == 1) {
        return "当前版本";
      }
      if (value == 2) {
        return "廉价版本";
      }
      return "";
    },
    parseYesOrNo: function(value) {
      if (value == 1) {
        return "是";
      }
      return "否";
    }
  },
  created() {
    this.kindTypeArray = kindTypeList;
    this.getList();
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
      dialogFormVisible: false,
      dialogContentVisible: false,
      kindTypeArray: [],
      temp: {
        sdkType: undefined,
        centerPlatformName: undefined
      },
      list: [],
      marryPrice: {
        price_1_1_1: undefined,
        price_1_1_2: undefined,
        price_1_1_3: undefined,
        price_1_4_1: undefined,
        price_1_4_2: undefined,
        price_1_4_3: undefined,
        price_2_1_1: undefined,
        price_2_1_2: undefined,
        price_2_1_3: undefined,
        price_2_4_1: undefined,
        price_2_4_2: undefined,
        price_2_4_3: undefined
      }
    };
  },
  methods: {
    handleFilter: function() {
      this.listQuery.pageIndex = 1;
      this.getList();
    },
    handleFlag: function(e) {
      this.dialogStatus = "create";
      this.dialogFormVisible = true;
      this.temp = Object.assign({}, e);
      console.log(this.temp);
    },
    updateflag: function() {
      updateCenterPlatformMarrySetFlag(this.temp).then(() => {
        this.getList();
        this.dialogFormVisible = false;
        this.showSuccess();
      });
    },
    handleContent: function(e) {
      this.dialogContentVisible = true;
      this.temp = Object.assign({}, e);
      console.log(this.temp.marrySet)
      this.marryPrice = this.getPriceFromArray(this.temp.marrySet)
    },
    updatecontent: function() {
      console.log(this.marryPrice);
      let marryArray = this.getArrayFromPrice(this.marryPrice);
      this.temp.marrySet = marryArray;
      console.log(this.temp.marrySet);
      updateCenterPlatformMarrySetContent(this.temp).then(() => {
        this.getList();
        this.dialogContentVisible = false;
        this.showSuccess();
      });
    },
    getList() {
      this.listLoading = true;
      let centerPlatformName = this.listQuery.centerPlatformName;
      let pageIndex = this.listQuery.pageIndex;
      getCenterPlatformMarrySetList(centerPlatformName, pageIndex)
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
    initMarryPrice() {
      let price = {
        price_1_1_1: undefined,
        price_1_1_2: undefined,
        price_1_1_3: undefined,
        price_1_4_1: undefined,
        price_1_4_2: undefined,
        price_1_4_3: undefined,
        price_2_1_1: undefined,
        price_2_1_2: undefined,
        price_2_1_3: undefined,
        price_2_4_1: undefined,
        price_2_4_2: undefined,
        price_2_4_3: undefined
      };
      return price;
    },
    getPriceFromArray(priceArray) {
      let price = this.initMarryPrice();
      
      let obj = eval('(' + priceArray + ')');
      for (let i = 0; i < obj.length; i++) {
        let item = obj[i];
        if (item.kindType == 1) {
          if (item.marryType == 1) {
            switch (item.marrySubType) {
              case 1:
                price.price_1_1_1 = item.useGold;
                break;
              case 2:
                price.price_1_1_2 = item.useGold;
                break;
              case 3:
                price.price_1_1_3 = item.useGold;
                break;
              default:
                break;
            }
          }
          if (item.marryType == 4) {
            switch (item.marrySubType) {
              case 1:
                price.price_1_4_1 = item.useGold;
                break;
              case 2:
                price.price_1_4_2 = item.useGold;
                break;
              case 3:
                price.price_1_4_3 = item.useGold;
                break;
              default:
                break;
            }
          }
        }
        if (item.kindType == 2) {
            if (item.marryType == 1) {
            switch (item.marrySubType) {
              case 1:
                price.price_2_1_1 = item.useGold;
                break;
              case 2:
                price.price_2_1_2 = item.useGold;
                break;
              case 3:
                price.price_2_1_3 = item.useGold;
                break;
              default:
                break;
            }
          }
          if (item.marryType == 4) {
            switch (item.marrySubType) {
              case 1:
                price.price_2_4_1 = item.useGold;
                break;
              case 2:
                price.price_2_4_2 = item.useGold;
                break;
              case 3:
                price.price_2_4_3 = item.useGold;
                break;
              default:
                break;
            }
          }
        }
      }
      return price;
    },
    getArrayFromPrice(price){
        let result = [];
        let item1 = {kindType:1,marryType:1,marrySubType:1,useGold:parseInt(price.price_1_1_1)}
        result.push(item1)
        let item2 = {kindType:1,marryType:1,marrySubType:2,useGold:parseInt(price.price_1_1_2)}
        result.push(item2)
        let item3 = {kindType:1,marryType:1,marrySubType:3,useGold:parseInt(price.price_1_1_3)}
        result.push(item3)
        let item4 = {kindType:1,marryType:4,marrySubType:1,useGold:parseInt(price.price_1_4_1)}
        result.push(item4)
        let item5 = {kindType:1,marryType:4,marrySubType:2,useGold:parseInt(price.price_1_4_2)}
        result.push(item5)
        let item6 = {kindType:1,marryType:4,marrySubType:3,useGold:parseInt(price.price_1_4_3)}
        result.push(item6)
        let item7 = {kindType:2,marryType:1,marrySubType:1,useGold:parseInt(price.price_2_1_1)}
        result.push(item7)
        let item8 = {kindType:2,marryType:1,marrySubType:2,useGold:parseInt(price.price_2_1_2)}
        result.push(item8)
        let item9 = {kindType:2,marryType:1,marrySubType:3,useGold:parseInt(price.price_2_1_3)}
        result.push(item9)
        let item10 = {kindType:2,marryType:4,marrySubType:1,useGold:parseInt(price.price_2_4_1)}
        result.push(item10)
        let item11 = {kindType:2,marryType:4,marrySubType:2,useGold:parseInt(price.price_2_4_2)}
        result.push(item11)
        let item12 = {kindType:2,marryType:4,marrySubType:3,useGold:parseInt(price.price_2_4_3)}
        result.push(item12)
        return result
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

