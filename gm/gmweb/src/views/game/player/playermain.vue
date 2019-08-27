<template>
    <div class="app-container">
        <el-tabs type="border-card">
            <el-tab-pane label="玩家邮件信息">
                <PlayerEMail :playerId="playerId" :serverId="serverId"></PlayerEMail>
            </el-tab-pane>
            <el-tab-pane label="物品变化">
                <PlayerItemChange :playerId="playerId" :serverId="serverId"></PlayerItemChange>
            </el-tab-pane>
            <el-tab-pane label="元宝变化">
                <PlayerGoldChange :playerId="playerId" :serverId="serverId"></PlayerGoldChange>
            </el-tab-pane>
            <el-tab-pane label="绑元变化">
                <PlayerBindgoldChange :playerId="playerId" :serverId="serverId"></PlayerBindgoldChange>
            </el-tab-pane>
            <el-tab-pane label="银两变化">
                <PlayerSilverChanged :playerId="playerId" :serverId="serverId"></PlayerSilverChanged>
            </el-tab-pane>
        </el-tabs>
    </div>
</template>
<script>
import waves from "@/directive/waves"; // 水波纹指令
import PlayerEMail from "./player_mail.vue";
import PlayerItemChange from "./player_item_change.vue";
import PlayerGoldChange from "./player_glod_changed.vue";
import PlayerBindgoldChange from "./player_bindgold_change.vue";
import PlayerSilverChanged from "./player_silver_changed.vue";
export default {
  name: "PlayerDetail",
  components: {
    PlayerEMail,
    PlayerItemChange,
    PlayerGoldChange,
    PlayerBindgoldChange,
    PlayerSilverChanged
  },
  directives: {
    waves
  },
  created() {
    const id = this.$route.params && this.$route.params.id;
    const serverId = this.$route.params && this.$route.params.serverId;
    this.playerId = id;
    this.serverId = serverId;
  },
  data() {
    return {
      playerId: undefined,
      serverId: undefined
    };
  },
  methods: {}
};
</script>

