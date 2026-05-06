<template>
  <section class="contact-edit">
    <header class="columns page-header">
      <div class="column is-10">
        <h1 class="title is-4 mb-2">
          <span v-if="isEditing">{{ form.firstName }} {{ form.lastName }}</span>
          <span v-else>New Contact</span>
        </h1>
      </div>
      <div class="column has-text-right">
        <b-button @click="goBack"> Back </b-button>
      </div>
    </header>

    <b-loading :active="loading.contact" />

    <div class="box">
      <form @submit.prevent="onSubmit">
        <b-field label="First Name" label-position="on-border">
          <b-input
            :maxlength="100"
            v-model="form.firstName"
            name="firstName"
            placeholder="First Name"
            required
          />
        </b-field>
        <b-field label="Last Name" label-position="on-border">
          <b-input
            :maxlength="100"
            v-model="form.lastName"
            name="lastName"
            placeholder="Last Name"
          />
        </b-field>
        <b-field label="Email" label-position="on-border">
          <b-input
            :maxlength="100"
            v-model="form.email"
            name="email"
            type="email"
            placeholder="Email"
          />
        </b-field>
        <b-field label="Phone" label-position="on-border">
          <b-input :maxlength="50" v-model="form.phone" name="phone" placeholder="Phone" />
        </b-field>

        <div class="field is-grouped is-grouped-right">
          <p class="control">
            <b-button @click="goBack">
              {{ $t('globals.buttons.cancel') }}
            </b-button>
          </p>
          <p class="control">
            <b-button
              native-type="submit"
              type="is-primary"
              :loading="loading.contacts"
              data-cy="btn-save"
            >
              {{ $t('globals.buttons.save') }}
            </b-button>
          </p>
        </div>
      </form>
    </div>
  </section>
</template>

<script>
import Vue from 'vue';
import { mapState } from 'vuex';

export default Vue.extend({
  name: 'ContactEdit',

  data() {
    return {
      isEditing: false,
      form: {
        id: null,
        firstName: '',
        lastName: '',
        email: '',
        phone: '',
      },
    };
  },

  computed: {
    ...mapState(['loading']),
  },

  mounted() {
    const { id } = this.$route.params;
    if (id && id !== 'new') {
      this.isEditing = true;
      this.fetchContact(id);
    }
  },

  methods: {
    fetchContact(id) {
      this.$api.getContact(id).then((response) => {
        this.form = { ...this.form, ...response };
      });
    },

    onSubmit() {
      if (this.isEditing) {
        this.updateContact();
      } else {
        this.createContact();
      }
    },

    createContact() {
      this.$api.createContact(this.form).then((data) => {
        this.$utils.toast(this.$t('globals.messages.created', { name: data.firstName }));
        this.$router.push({ name: 'contact', params: { id: data.id } });
      });
    },

    updateContact() {
      this.$api.updateContact(this.form).then((data) => {
        this.$utils.toast(this.$t('globals.messages.updated', { name: data.firstName }));
        this.$router.push({ name: 'contact', params: { id: data.id } });
      });
    },

    goBack() {
      if (this.isEditing) {
        this.$router.push({ name: 'contact', params: { id: this.form.id } });
      } else {
        this.$router.push({ name: 'contacts' });
      }
    },
  },
});
</script>

<style lang="scss" scoped>
.contact-edit {
  .box {
    max-width: 800px;
    margin: 0 auto;
  }
}
</style>
