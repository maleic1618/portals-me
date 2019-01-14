<template>
  <v-container grid-list-md fluid>
    <v-flex xs12 class="collection-title">
      <v-layout row wrap>
        <v-flex xs5>
          <h2><v-icon>collections</v-icon> myuon / myuon</h2>
        </v-flex>
        <v-spacer />
        <v-btn
          dark
          depressed
          outline
          color="indigo"
          @click="createArticleDialog = true"
        >
          <v-icon left>add</v-icon>
          作品を登録
        </v-btn>

        <v-dialog
          v-model="createArticleDialog"
          max-width="760px"
        >
          <v-card>
            <v-card-title
              class="headline grey lighten-2"
              primary-title
            >
              作品を登録
            </v-card-title>

            <v-card-text>
              <v-container fluid>
                <div ref="oEmbedPreview">
                  ここにプレビューが表示されます…
                </div>
              </v-container>

              <v-form
                v-model="createArticleForm.valid"
                ref="createArticleForm"
                lazy-validation
              >
                <v-text-field
                  v-model="createArticleForm.url"
                  label="URLを入力"
                  :rules="[v => !!v || '必須項目です']"
                  required
                  @input="previewOEmbed"
                />

                <v-text-field
                  v-model="createArticleForm.title"
                  label="タイトル"
                  :rules="[v => !!v || '必須項目です']"
                  required
                />

                <v-textarea
                  v-model="createArticleForm.description"
                  label="説明(任意)"
                  auto-grow
                  rows="1"
                />

                <v-checkbox
                  v-model="createArticleForm.checkbox"
                  :rules="[v => !!v || '作品を登録できるのは正当な権利者のみです']"
                  label="私はこの作品の正当な権利者であり、他のいかなる権利の侵害もしていません"
                  required
                >
                </v-checkbox>

                <v-btn
                  :disabled="!createArticleForm.valid"
                  color="indigo"
                  dark
                  @click="submit"
                >
                  <v-icon left>send</v-icon>
                  送信
                </v-btn>
              </v-form>
            </v-card-text>
          </v-card>
        </v-dialog>
      </v-layout>
    </v-flex>

    <v-tabs
      class="collection-tab"
    >
      <v-tab>
        <v-icon>layers</v-icon>
        Cards
      </v-tab>

      <v-tab>
        <v-icon>chat</v-icon>
        Channel
      </v-tab>

      <v-tab>
        <v-icon>settings</v-icon>
        Settings
      </v-tab>

      <v-tab-item class="collection-layout">
        <v-layout flex-child wrap>
          <v-hover>
            <v-flex
              md3
              d-flex
              slot-scope="{ hover }"
              @click="clickArticleCard"
            >
              <v-card
                class="mx-auto"
                :class="`elevation-${hover ? 6 : 2}`"
              >
                <v-img
                  :aspect-ratio="16/9"
                  style="background-color: #eeeeee;"
                >
                </v-img>
                <v-card-title>
                  Project Meow
                </v-card-title>
              </v-card>
            </v-flex>
          </v-hover>

          <v-hover>
            <v-flex
              md3
              d-flex
              slot-scope="{ hover }"
              @click="clickArticleCard"
            >
              <v-card
                class="mx-auto"
                :class="`elevation-${hover ? 6 : 2}`"
              >
                <v-img
                  :aspect-ratio="16/9"
                  style="background-color: #80CBC4;"
                >
                </v-img>
                <v-card-title>
                  Project Meow
                </v-card-title>
              </v-card>
            </v-flex>
          </v-hover>

          <v-hover>
            <v-flex
              md3
              d-flex
              slot-scope="{ hover }"
              @click="clickArticleCard"
            >
              <v-card
                class="mx-auto"
                :class="`elevation-${hover ? 6 : 2}`"
              >
                <v-img
                  :aspect-ratio="16/9"
                  style="background-color: #EF9A9A;"
                >
                </v-img>
                <v-card-title>
                  Project Meow
                </v-card-title>
              </v-card>
            </v-flex>
          </v-hover>

          <v-hover>
            <v-flex
              md3
              d-flex
              slot-scope="{ hover }"
              @click="clickArticleCard"
            >
              <v-card
                class="mx-auto"
                :class="`elevation-${hover ? 6 : 2}`"
              >
                <v-img
                  :aspect-ratio="16/9"
                  style="background-color: #9FA8DA;"
                >
                </v-img>
                <v-card-title>
                  Project Meow
                </v-card-title>
              </v-card>
            </v-flex>
          </v-hover>
        </v-layout>

        <v-dialog
          v-model="articleDialog"
          max-width="600"
        >
          <v-card>
            <v-card-title class="headline">Article #0</v-card-title>

            <v-card-text>
              ほげほげ
            </v-card-text>
          </v-card>
        </v-dialog>
      </v-tab-item>

      <v-tab-item>
        <v-layout>
          <v-flex xs12 class="message">
            <v-avatar>
              <v-img
                src="https://lh6.googleusercontent.com/-HrqEjsNu_No/AAAAAAAAAAI/AAAAAAAAAMI/Rg4RwE9Y7So/s96-c/photo.jpg"
              />
            </v-avatar>

            <div class="content">
              <div class="header">
                <strong>myuon</strong>
              </div>

              <div class="input-area">
                <div>
                  <autogrow-textarea placeholder="myuon/myuonへのメッセージ…" />
                </div>

                <v-btn depressed color="primary">
                  送信
                </v-btn>
              </div>
            </div>
          </v-flex>
        </v-layout>

        <v-divider />

        <v-layout>
          <v-flex xs12 class="message">
            <v-avatar>
              <v-img
                src="https://lh6.googleusercontent.com/-HrqEjsNu_No/AAAAAAAAAAI/AAAAAAAAAMI/Rg4RwE9Y7So/s96-c/photo.jpg"
              />
            </v-avatar>

            <div class="content">
              <div class="header">
                <strong>myuon</strong>

                10分前
              </div>

              <div class="content-content">
                This is 最高にちょうどいい本文。
              </div>
            </div>
          </v-flex>
        </v-layout>

        <v-layout>
          <v-flex xs12 class="message">
            <v-avatar>
              <v-img
                src="https://lh6.googleusercontent.com/-HrqEjsNu_No/AAAAAAAAAAI/AAAAAAAAAMI/Rg4RwE9Y7So/s96-c/photo.jpg"
              />
            </v-avatar>

            <div class="content">
              <div class="header">
                <strong>myuon</strong>

                10分前
              </div>

              <div class="content-content">
                This is 最高にちょうどいい本文。This is 最高にちょうどいい本文。This is 最高にちょうどいい本文。This is 最高にちょうどいい本文。This is 最高にちょうどいい本文。This is 最高にちょうどいい本文。This is 最高にちょうどいい本文。This is 最高にちょうどいい本文。This is 最高にちょうどいい本文。This is 最高にちょうどいい本文。This is 最高にちょうどいい本文。This is 最高にちょうどいい本文。This is 最高にちょうどいい本文。This is 最高にちょうどいい本文。This is 最高にちょうどいい本文。This is 最高にちょうどいい本文。This is 最高にちょうどいい本文。This is 最高にちょうどいい本文。This is 最高にちょうどいい本文。This is 最高にちょうどいい本文。This is 最高にちょうどいい本文。This is 最高にちょうどいい本文。This is 最高にちょうどいい本文。This is 最高にちょうどいい本文。This is 最高にちょうどいい本文。This is 最高にちょうどいい本文。This is 最高にちょうどいい本文。This is 最高にちょうどいい本文。This is 最高にちょうどいい本文。This is 最高にちょうどいい本文。This is 最高にちょうどいい本文。
              </div>
            </div>
          </v-flex>
        </v-layout>

        <v-layout>
          <v-flex xs12 class="message">
            <v-avatar>
              <v-img
                src="https://lh6.googleusercontent.com/-HrqEjsNu_No/AAAAAAAAAAI/AAAAAAAAAMI/Rg4RwE9Y7So/s96-c/photo.jpg"
              />
            </v-avatar>

            <div class="content">
              <div class="header">
                <strong>myuon</strong>

                10分前
              </div>

              <div class="content-content">
                L <br />
                o <br />
                o <br />
                o <br />
                o <br />
                o <br />
                o <br />
                o <br />
                o <br />
                o <br />
                o <br />
                o <br />
                o <br />
                o <br />
                o <br />
                o <br />
                o <br />
                g cat <br />
              </div>
            </div>
          </v-flex>
        </v-layout>
      </v-tab-item>

      <v-tab-item>
        <v-container>
          <v-btn
            color="error"
          >
            プロジェクトを削除
            <v-icon right dark>delete</v-icon>
          </v-btn>
        </v-container>
      </v-tab-item>
    </v-tabs>
  </v-container>
</template>

<script>
import AutogrowTextarea from '@/components/AutogrowTextarea';
import fetchJsonp from 'fetch-jsonp';

export default {
  data () {
    return {
      articleDialog: false,
      createArticleDialog: false,
      createArticleDialogTab: null,
      createArticleForm: {
        valid: true,
        url: '',
        checkbox: false,
        title: '',
        description: '',
      },
    };
  },
  components: {
    SheetFooter: {
      functional: true,

      render (h, { children }) {
        return h('v-sheet', {
          staticClass: 'mt-auto align-center justify-center d-flex',
          props: {
            color: 'rgba(0, 0, 0, .36)',
            dark: true,
            height: 50
          }
        }, children)
      }
    },
    AutogrowTextarea,
  },
  methods: {
    clickArticleCard () {
      this.articleDialog = true;
    },
    submit () {
      if (this.$refs.createArticleForm.validate()) {
        console.log(this.createArticleForm);
      }
    },
    async previewOEmbed () {
      const getProvider = (url) => {
        if (/https:\/\/twitter\.com\/.*\/status\/.*/.test(url)) {
          console.log('twitter!');
          return `https://publish.twitter.com/oembed?format=json&url=${encodeURIComponent(url)}`
        }
      };
      const area = this.$refs.oEmbedPreview;

      const url = getProvider(this.createArticleForm.url);
      console.log(url);

      if (!url) {
        area.innerText = 'ここにプレビューが表示されます…';
        return;
      };

      const response = await fetchJsonp(url);
      const card_json = await response.json();
      console.log(card_json);

      const replaceHTML = (element, html) => {
        element.innerHTML = html;
        element.querySelectorAll('script').forEach(scriptElement => {
          const se = document.createElement('script');
          se.src = scriptElement.src;
          console.log(scriptElement.src);
          scriptElement.replaceWith(se);
        });
      }

      replaceHTML(area, card_json.html);
    },
  }
}
</script>

<style scoped>
.collection-title {
  margin-bottom: 1em;
}

.collection-tab .v-icon {
  margin-right: 0.2em;
}

.collection-layout .card {
  margin-top: 1rem;
}

.message {
  display: flex;
}

.message .v-avatar {
  display: block;
  flex: 0 0 auto;
  height: auto;
  margin-right: 20px;
}

.message .content {
  display: block;
  flex: 1 1 auto;
  margin-top: 1em;
}

.message .input-area {
  display: flex;
}

.message .input-area > div {
  display: block;
  flex: 1 1 auto;
}

.message .input-area > .v-btn {
  margin-top: 0;
  margin-bottom: 0;
}

.message .content .header {
  margin-bottom: 0.3em;
}
</style>