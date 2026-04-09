<template>
  <form @submit.prevent="onSubmit">
    <div class="modal-card content" style="width: auto">
      <header class="modal-card-head">
        <h4 class="modal-card-title">{{ isEditing ? 'Edit Contact' : 'Create Contact' }}</h4>
      </header>

      <section class="modal-card-body">
        <b-field label="First Name" label-position="on-border">
          <b-input v-model="form.firstName" required />
        </b-field>
        <b-field label="Last Name" label-position="on-border">
          <b-input v-model="form.lastName" required />
        </b-field>
        <b-field label="Email" label-position="on-border">
          <b-input v-model="form.email" required />
        </b-field>
        <b-field label="Phone" label-position="on-border">
          <b-input v-model="form.phone" required />
        </b-field>
        <b-field label="Contact Type" label-position="on-border">
          <b-select v-model="form.contactType" required>
            <option value="lead">Lead</option>
            <option value="customer">Customer</option>
          </b-select>
        </b-field>
        <b-field label="Status" label-position="on-border">
          <b-select v-model="form.status" required>
            <option value="active">Active</option>
            <option value="inactive">Inactive</option>
          </b-select>
        </b-field>
        <b-field label="Description" label-position="on-border">
          <b-input
            :maxlength="2000"
            v-model="form.description"
            name="description"
            :placeholder="$t('globals.fields.description')"
            type="textarea"
          />
        </b-field>
      </section>

      <footer class="modal-card-foot">
        <b-button @click="$parent.close()">
          {{ $t('globals.buttons.close') }}
        </b-button>
        <b-button
          v-if="$can('contacts:manager') || $canList(data.id, 'contacts:manage')"
          native-type="submit"
          type="is-primary"
          :loading="loading.contacts"
          data-cy="btn-save"
        >
          {{ $t('globals.buttons.save') }}
        </b-button>
      </footer>
    </div>
  </form>
</template>

<script>
import Vue from 'vue';
import { mapState } from 'vuex';

export default Vue.extend({
  name: 'ContactForm',

  props: {
    data: { type: Object, default: () => ({}) },
    isEditing: { type: Boolean, default: false },
  },

  data() {
    return {
      form: {
        firstName: '',
        lastName: '',
        email: '',
        phone: '',
        contactType: '',
        status: '',
        description: '',
      },
    };
  },

  computed: {
    ...mapState(['loading']),
  },

  methods: {
    onSubmit() {
      if (this.isEditing) {
        this.updateContact();
      } else {
        this.createContact();
      }
    },

    createContact() {
      return this.$api.createContact(this.form).then((data) => {
        this.$emit('finished');
        this.$parent.close();
        this.$utils.toast(
          // eslint-disable-next-line comma-dangle
          this.$t('globals.messages.created', { name: `${data.firstName} ${data.lastName}` })
        );
      });
    },

    updateContact() {
      return this.$api.updateContact(this.form).then((data) => {
        this.$emit('finished');
        this.$parent.close();
        this.$utils.toast(
          // eslint-disable-next-line comma-dangle
          this.$t('globals.messages.updated', { name: `${data.firstName} ${data.lastName}` })
        );
      });
    },
  },

  mounted() {
    this.form = { ...this.form, ...this.$props.data };
  },
});
</script>
