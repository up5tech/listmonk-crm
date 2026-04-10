<template>
  <section class="contact-detail">
    <header class="columns page-header">
      <div class="column is-10">
        <h1 class="title is-4 mb-2">{{ title }}</h1>
      </div>
    </header>

    <b-loading :active="loading.contact" />

    <b-tabs type="is-boxed" :animated="false" v-model="activeTab" @input="onTab">
      <b-tab-item
        label="Information"
        label-position="on-border"
        value="info"
        icon="rocket-launch-outline"
      >
        <section class="wrap">
          <div class="columns">
            <div class="column is-6">
              <strong>First Name</strong>:
              <span>{{ contact.firstName || '' }}</span>
            </div>
            <div class="column is-6">
              <strong>Last Name</strong>:
              <span>{{ contact.lastName || '' }}</span>
            </div>
            <div class="column is-6">
              <strong>Email</strong>:
              <span>{{ contact.email || '' }}</span>
            </div>
          </div>
          <div class="columns">
            <div class="column is-6">
              <strong>Phone</strong>:
              <span>{{ contact.phone || '' }}</span>
            </div>
          </div>
        </section>
      </b-tab-item>
    </b-tabs>
  </section>
</template>

<script>
import Vue from 'vue';
import { mapState } from 'vuex';

export default Vue.extend({
  name: 'ContactDetail',

  data() {
    return {
      title: 'Contact Detail',
      contact: {},
      activeTab: 'info',
    };
  },

  methods: {
    fetchContact(id) {
      this.$api.getContact(id).then((response) => {
        this.contact = response || {};
        if (response.firstName || response.lastName) {
          this.title = `${response.firstName || ''} ${response.lastName || ''}`.trim();
        } else {
          this.title = 'Contact Detail';
        }
      });
    },

    onTab(tab) {
      this.activeTab = tab;
    },
  },

  computed: {
    ...mapState(['loading', 'settings']),
  },

  mounted() {
    const { id } = this.$route.params;
    if (id) {
      this.fetchContact(id);
    }
  },
});
</script>
