<template>
  <section class="contact-detail">
    <h1 class="title is-4 mb-2">{{ title }}</h1>
  </section>
</template>

<script>
import Vue from 'vue';

export default Vue.extend({
  name: 'ContactDetail',

  data() {
    return {
      title: 'Contact Detail',
      contact: {},
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
  },

  mounted() {
    const { id } = this.$route.params;
    if (id) {
      this.fetchContact(id);
    }
  },
});
</script>
