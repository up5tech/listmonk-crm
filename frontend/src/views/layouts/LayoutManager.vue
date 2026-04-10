<template>
  <div>
    <h1 class="text-2xl font-bold">Layout Manager</h1>

    <div class="flex">
      <div class="w-[280px] mt-2">
        <ul class="unordered-list">
          <li v-for="module in modules" :key="module.name" class="mb-2">
            <div class="font-bold">
              {{ module.label }}
            </div>
            <ul class="unordered-list mt-2 ml-8">
              <li v-for="layoutType in layoutTypes" :key="layoutType" class="mb-2">
                <button
                  type="button"
                  class="w-full border border-solid border-gray-300 rounded-md px-4 py-2 text-sm text-left cursor-pointer hover:bg-gray-100"
                  @click="getLayouts(module.name, layoutType)"
                  v-if="module.name !== currentModule || layoutType !== currentLayoutType"
                >
                  {{ layoutType }}
                </button>
                <button
                  type="button"
                  class="w-full border border-solid border-gray-300 rounded-md px-4 py-2 text-sm text-left font-bold bg-gray-100"
                  v-else
                >
                  {{ layoutType }}
                </button>
              </li>
            </ul>
          </li>
        </ul>
      </div>
    </div>
  </div>
</template>

<script>
import Vue from 'vue';

export default Vue.extend({
  name: 'LayoutManager',

  data() {
    return {
      modules: [],
      layouts: {},
      layoutTypes: ['list', 'record'],
      currentModule: '',
      currentLayoutType: '',
    };
  },

  methods: {
    getModules() {
      this.$api.getModules().then((response) => {
        this.modules = response;
      });
    },

    getLayouts(module, layoutType) {
      this.currentModule = module;
      this.currentLayoutType = layoutType;
      this.$api.getLayouts(module, layoutType).then((response) => {
        this.layouts = response;
      });
    },
  },

  mounted() {
    this.getModules();
  },
});
</script>

<style lang="scss">
@import '../../assets/tailwind.css';
</style>
