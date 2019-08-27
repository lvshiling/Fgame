<template>
    <div class="app-container">
        <el-tabs type="border-card">
            <el-tab-pane label="单服提现记录">
                <GameFeedBackFeeList></GameFeedBackFeeList>
            </el-tab-pane>
            <el-tab-pane label="中心服提现记录">
              <CenterFeedBackFeeList></CenterFeedBackFeeList>
            </el-tab-pane>
        </el-tabs>
    </div>
</template>
<script>
import waves from "@/directive/waves"; // 水波纹指令
import GameFeedBackFeeList from './gamefeedbackfeelist.vue'
import CenterFeedBackFeeList from './centerfeedbackfeelist.vue'
export default {
  name: "FeedBackFee",
  components: {
    GameFeedBackFeeList,
    CenterFeedBackFeeList
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

