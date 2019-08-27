<template>
    <div class="app-container">
        <el-tabs type="border-card">
            <el-tab-pane label="单服充值记录">
                <GameOrderList></GameOrderList>
            </el-tab-pane>
            <el-tab-pane label="单服汇总">
              <GameOrderStatic></GameOrderStatic>
            </el-tab-pane>
            <el-tab-pane label="中心服充值记录">
              <CenterOrderList></CenterOrderList>
            </el-tab-pane>
            <el-tab-pane label="中心充值汇总">
              <CenterOrderStatic></CenterOrderStatic>
            </el-tab-pane>
            <el-tab-pane label="多服汇总">
              <CenterOrderDateStatic></CenterOrderDateStatic>
            </el-tab-pane>
            <el-tab-pane label="充值汇总">
              <OrderTotalList></OrderTotalList>
            </el-tab-pane>
        </el-tabs>
    </div>
</template>
<script>
import waves from "@/directive/waves"; // 水波纹指令
import CenterOrderList from './ordestatic/centerorderlist.vue'
import GameOrderList from './ordestatic/orderlist.vue'
import GameOrderStatic from './ordestatic/gameorderstatic.vue'
import CenterOrderStatic from './ordestatic/centerorderstatic.vue'
import CenterOrderDateStatic from './ordestatic/centerorderdate.vue'
import OrderTotalList from './ordestatic/ordertotallist'
export default {
  name: "OrderStatic",
  components: {
    CenterOrderList,
    GameOrderList,
    GameOrderStatic,
    CenterOrderStatic,
    CenterOrderDateStatic,
    OrderTotalList
  },
  directives: {
    waves
  },
  created() {},
  data() {
    return {};
  },
  methods: {}
};
</script>

