<template>
  <form @submit.prevent="onSubmit">
    <div class="modal-card content" style="width: auto">
      <header class="modal-card-head">
        <h4 class="modal-card-title">{{ $t('contacts.form.title') }}</h4>
      </header>
      <section class="modal-card-body">
        <b-field :label="$t('contacts.form.fields.name')">
          <b-input v-model="form.name" required />
        </b-field>
        <b-field :label="$t('contacts.form.fields.email')">
          <b-input v-model="form.email" required />
        </b-field>
        <b-field :label="$t('contacts.form.fields.phone')">
          <b-input v-model="form.phone" required />
        </b-field>
      </section>
    </div>
  </form>
</template>

<script>
import Vue from 'vue';
import { mapState } from 'vuex';

export default Vue.extend({
  name: 'ContactForm',

  data() {
    return {
      form: {
        name: '',
        email: '',
        phone: '',
      },
    };
  },

  computed: {
    ...mapState(['loading']),
  },

  methods: {
    onSubmit() {
      this.$api.createContact(this.form).then((data) => {
        this.$emit('finished');
        this.$parent.close();
        this.$utils.toast(this.$t('globals.messages.created', { name: data.name }));
      });
    },
  },
});
</script>
