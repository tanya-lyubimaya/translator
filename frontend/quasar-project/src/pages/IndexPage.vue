<template>
  <q-layout view="hHh lpR fFf">
    <q-header reveal elevated class="bg-primary text-white">
      <q-toolbar>
        <q-toolbar-title>
          <q-avatar>
            <img src="https://cdn.quasar.dev/logo-v2/svg/logo-mono-white.svg" />
          </q-avatar>
          Title
        </q-toolbar-title>
      </q-toolbar>
    </q-header>

    <q-page-container>
      <router-view />
      <div class="fit row wrap justify-center items-center content-start">
        <q-select
          v-model="inputLanguage"
          use-input
          input-debounce="0"
          label="Language"
          :options="languageOptions"
          @filter="filterFn"
          style="width: 250px"
          behavior="menu"
        />
        <q-select
          v-model="outputLanguage"
          use-input
          input-debounce="0"
          label="Language"
          :options="languages"
          style="width: 250px"
          behavior="menu"
          class="col-4 offset-1"
        />
      </div>
      <div class="fit row wrap justify-center content-start">
        <q-input
          autogrow
          outlined
          clearable
          use-input
          input-debounce="0"
          color="orange"
          v-model="input"
          type="textarea"
          label="Text"
          class="col-4"
          style="overflow: auto"
          clear-icon="close"
        />
        <q-input
          autogrow
          outlined
          disable
          color="orange"
          v-model="output"
          type="textarea"
          label="Translation"
          class="col-4 offset-1"
          style="overflow: auto"
        />
      </div>
      <q-btn
        push
        color="white"
        text-color="primary"
        label="Push"
        @click="sendForTranslation"
      />
    </q-page-container>
  </q-layout>
</template>

<script>
import { ref } from "vue";

export default {
  setup() {
    return {
      output: ref(null),
      inputLanguage: ref("ru"),
      outputLanguage: ref("en"),
    };
  },
  data() {
    return { languages: [], input: "", languageOptions: [] };
  },
  beforeMount() {
    const path = "http://localhost:8080/get-languages";
    this.$axios.get(path).then(
      (res) => {
        this.languages = res.data;
        this.languageOptions = this.languages;
        console.log(res);
      },
      (error) => {
        console.error(error);
      }
    );
  },
  methods: {
    sendForTranslation() {
      const url = "http://localhost:8080/get-translation";
      this.$axios
        .post(url, {
          source_lang: this.inputLanguage,
          target_lang: this.outputLanguage,
          source_text: this.input,
        })
        .catch((err) => {
          this.$q.notify({
            position: this.notificationsPos,
            icon: "warning",
            type: "negative",
            multiLine: true,
            message: "Возникла ошибка!",
          });
        });
    },
    filterFn(val, update) {
      if (val === "") {
        update(() => {
          this.languageOptions.value = this.languages;

          // here you have access to "ref" which
          // is the Vue reference of the QSelect
        });
        return;
      }

      update(() => {
        const needle = val.toLowerCase();
        this.languageOptions.value = this.languages.filter(
          (v) => v.toLowerCase().indexOf(needle) > -1
        );
      });
    },
  },
};
</script>
