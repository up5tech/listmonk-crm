<template>
  <section class="account-edit">
    <header class="columns page-header">
      <div class="column is-10">
        <h1 class="title is-4 mb-2">
          <span v-if="isEditing">{{ form.name }}</span>
          <span v-else>New Account</span>
        </h1>
      </div>
      <div class="column has-text-right">
        <b-button @click="goBack"> Back </b-button>
      </div>
    </header>

    <b-loading :active="loading.account" />

    <div class="box">
      <form @submit.prevent="onSubmit">
        <b-field :label="$t('globals.fields.name')" label-position="on-border">
          <b-input
            :maxlength="200"
            :ref="'focus'"
            v-model="form.name"
            name="name"
            :placeholder="$t('globals.fields.name')"
            required
          />
        </b-field>
        <b-field label="Short Name" label-position="on-border">
          <b-input
            :maxlength="100"
            v-model="form.shortName"
            name="shortName"
            placeholder="Short Name"
          />
        </b-field>
        <b-field label="Account Type" label-position="on-border">
          <b-select v-model="form.accountType" name="accountType">
            <option value="direct">Direct</option>
            <option value="agency">Agency</option>
          </b-select>
        </b-field>
        <b-field label="Industry" label-position="on-border">
          <b-select v-model="form.industry" name="industry">
            <option value="industry">Industry</option>
          </b-select>
        </b-field>
        <b-field label="Sic Code" label-position="on-border">
          <b-input :maxlength="10" v-model="form.sicCode" name="sicCode" placeholder="Sic Code" />
        </b-field>
        <b-field label="Website" label-position="on-border">
          <b-input :maxlength="100" v-model="form.website" name="website" placeholder="Website" />
        </b-field>
        <b-field :label="$t('globals.fields.description')" label-position="on-border">
          <b-input
            :maxlength="2000"
            v-model="form.description"
            name="description"
            :placeholder="$t('globals.fields.description')"
            type="textarea"
          />
        </b-field>

        <div class="field is-grouped is-grouped-right">
          <p class="control">
            <b-button @click="goBack">
              {{ $t('globals.buttons.cancel') }}
            </b-button>
          </p>
          <p class="control">
            <b-button
              v-if="$can('accounts:manager') || $canList(form.id, 'accounts:manage')"
              native-type="submit"
              type="is-primary"
              :loading="loading.accounts"
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
  name: 'AccountEdit',

  data() {
    return {
      isEditing: false,
      form: {
        id: null,
        name: '',
        shortName: '',
        accountType: '',
        industry: '',
        sicCode: '',
        website: '',
        description: '',
      },
    };
  },

  computed: {
    ...mapState(['loading', 'profile']),
  },

  mounted() {
    const { id } = this.$route.params;
    if (id && id !== 'new') {
      this.isEditing = true;
      this.fetchAccount(id);
    }
  },

  methods: {
    fetchAccount(id) {
      this.$api.getAccount(id).then((response) => {
        this.form = { ...this.form, ...response };
      });
    },

    onSubmit() {
      if (this.isEditing) {
        this.updateAccount();
      } else {
        this.createAccount();
      }
    },

    createAccount() {
      this.$api.createAccount(this.form).then((data) => {
        this.$utils.toast(this.$t('globals.messages.created', { name: data.name }));
        this.$router.push({ name: 'account', params: { id: data.id } });
      });
    },

    updateAccount() {
      this.$api.updateAccount(this.form).then((data) => {
        this.$utils.toast(this.$t('globals.messages.updated', { name: data.name }));
        this.$router.push({ name: 'account', params: { id: data.id } });
      });
    },

    goBack() {
      if (this.isEditing) {
        this.$router.push({ name: 'account', params: { id: this.form.id } });
      } else {
        this.$router.push({ name: 'accounts' });
      }
    },
  },
});
</script>

<style lang="scss" scoped>
.account-edit {
  .box {
    max-width: 800px;
    margin: 0 auto;
  }
}
</style>
