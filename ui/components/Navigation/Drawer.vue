<template>
  <v-navigation-drawer
    app
    :permanent="$vuetify.breakpoint.mdAndUp"
    hide-overlay
    v-model="show"
  ></v-navigation-drawer>
</template>

<script>
export default {
  props: ["drawer"],
  model: {
    prop: "drawer",
    event: "drawerChange",
  },
  watch: {
    drawer: function (val) {
      this.show = val;
    },
    show: function (val) {
      if (!val) {
        this.$emit("update:drawer", val);
      }
    },
    data() {
      return {
        show: this.drawer,
      };
    },
  },
};
</script>

<style>
</style>