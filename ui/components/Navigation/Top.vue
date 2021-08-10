<template>
  <v-app-bar
    app
    flat
    style="border-bottom: solid 1.5px #e0e0e0; background-color: white"
  >
    <v-spacer></v-spacer>
    <v-app-bar-nav-icon
      @click="toggleDrawer()"
      v-if="!vuetify.breakpoint.mdAndUp"
    ></v-app-bar-nav-icon>
  </v-app-bar>
</template>

<script>
export default {
  props: ["drawer"],
  model: {
    prop: "drawer",
    event: "drawerChange",
  },
  data() {
    return {};
  },
  computed: {
    toggleDrawer() {
      this.$emit("update:drawer", !this.drawer);
    },
  },
  methods: {},
};
</script>

<style>
</style>