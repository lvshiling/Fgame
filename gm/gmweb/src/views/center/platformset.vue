<template>
    <div class="app-container">
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
            <el-table-column label="中心平台名称" align="center" width="120px">
                <template slot-scope="scope">
                    <span>{{ scope.row.centerPlatformName }}</span>
                </template>
            </el-table-column>
            <el-table-column  v-for="(item,index) in metaList" :key="index" :label="item.tipName" width="120px" align="left">
                <template slot-scope="scope">
                    <span>{{ scope.row.settingObj | parseSettingItem(item.code,item)}}</span>
                </template>
            </el-table-column>
            <el-table-column label="操作" align="center" width="280px" class-name="small-padding fixed-width">
                <template slot-scope="scope">
                <el-button type="primary" size="mini" @click="handleFlag(scope.row)">设置</el-button>
                </template>
            </el-table-column>
        </el-table>

        <div class="pagination-container" style="margin-top:15px;">
            <el-pagination :current-page="listQuery.pageIndex" :page-sizes="[20]" :total="total" background layout="total, sizes, prev, pager, next, jumper"  @current-change="handleCurrentChange"/>
        </div>

        <el-dialog :title="textMap[dialogStatus]" :visible.sync="dialogFormVisible">
        <el-form ref="dataForm" :model="temp" label-position="left" label-width="120px" style="width: 400px; margin-left:50px;">
            <el-form-item label="中心平台名">
                <el-input v-model="temp.centerPlatformId" disabled/>
            </el-form-item>
            
            <el-form-item  v-for="(item,index) in metaList" :key="index" :label="item.tipName" >
                <el-input v-model="tempSetting[item.code]" v-if="item.settingType == 1"/>
                <el-input-number v-model="tempSetting[item.code]" v-if="item.settingType == 2"/>
                <el-select v-model="tempSetting[item.code]" :placeholder="item.tipName" style="width: 160px" class="filter-item" v-if="item.settingType == 3">
                    <el-option v-for="selectItem in item.selectItem" :key="selectItem.key" :label="selectItem.value" :value="selectItem.key" />
                </el-select>
            </el-form-item>
        </el-form>
        <div slot="footer" class="dialog-footer">
            <el-button @click="dialogFormVisible = false">取消</el-button>
            <el-button type="primary" @click="updateflag">确定</el-button>
        </div>
        </el-dialog>
    </div>
</template>

<script>
import waves from "@/directive/waves"; // 水波纹指令
import {
  getCenterPlatformSetList,
  updateCenterPlatformSaveSetting,
  getCenterPlatformMetaSetting
} from "@/api/centerPlatform";
import { refreshCenterServer } from "@/api/centerserver";
import { kindTypeList } from "@/types/center";
import { getAllSdkType } from "@/api/center";
import { Switch } from "element-ui";
export default {
  name: "centerPlatformSettingList",
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
    parseYesOrNo: function(value) {
      if (value == 1) {
        return "是";
      }
      return "否";
    },
    parseSettingItem: function(value, code, metaItem) {
      if (!value) {
        return "";
      }
      if (!metaItem) {
        return value;
      }
      let rstValue = undefined;
      if (value.hasOwnProperty(code)) {
        rstValue = value[code];
      }
      if (metaItem.settingType == 3) {
        //选择框
        for (let i = 0; i < metaItem.selectItem.length; i++) {
          let item = metaItem.selectItem[i];
          if(item.key == rstValue){
            return item.value
          }
        }
      }
      return rstValue;
    }
  },
  created() {
    this.kindTypeArray = kindTypeList;
    this.getList();
    this.getMeta();
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
        centerPlatformId: undefined
      },
      list: [],
      metaList: [],
      tempSetting: {}
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
      this.tempSetting = this.jsonToObject(this.temp.setting);
      if (!this.tempSetting) {
        this.tempSetting = {};
        for (let i = 0; i < this.metaList.length; i++) {
          let code = this.metaList[i].code;
          this.tempSetting[code] = undefined;
        }
      }
    },
    updateflag: function() {
      this.temp.setting = JSON.stringify(this.tempSetting);
      updateCenterPlatformSaveSetting(this.temp).then(() => {
        this.getList();
        this.dialogFormVisible = false;
        this.showSuccess();
      });
    },
    getList() {
      this.listLoading = true;
      let centerPlatformName = this.listQuery.centerPlatformName;
      let pageIndex = this.listQuery.pageIndex;
      getCenterPlatformSetList(centerPlatformName, pageIndex)
        .then(res => {
          this.list = res.itemArray;
          this.total = res.total;
          this.listLoading = false;
          for (let i = 0; i < this.list.length; i++) {
            this.list[i].settingObj = this.jsonToObject(this.list[i].setting);
          }
        })
        .catch(() => {
          this.listLoading = false;
        });
    },
    handleCurrentChange(e) {
      this.listQuery.pageIndex = e;
      this.getList();
    },
    getMeta() {
      getCenterPlatformMetaSetting()
        .then(res => {
          this.metaList = res.itemArray;
          this.listLoading = false;
        })
        .catch(() => {
          this.listLoading = false;
        });
    },
    jsonToObject(jsonString) {
      if (jsonString) {
        return JSON.parse(jsonString);
      }
      let defaultObj = {};
      return defaultObj;
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

