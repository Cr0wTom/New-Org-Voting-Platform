import {AjaxCaller} from '../modules/ajaxCaller.js';

let topUrls = new Vue({
        el: '#topUrlsSection',
        delimiters: ['${', '}'],
        data: {
            loaderVisible: false,

            language: 'en',
            collection: 'terms',
            urls: []
        },
        methods: {
            initTopUrls() {
                this.loaderVisible = true;
                this.updateTopUrls();
            }
            ,
            updateTopUrls: async function () {

                let endpoint = '/api/tweets/top-urls';
                let data = {
                    by: this.by,
                    language: this.language,
                    collection: this.collection,
                    num: this.num
                };
                let response = await AjaxCaller.makeInternalCall(endpoint, data);
                this.createEmbeddedUrls(response.data);

                this.loaderVisible = false;
            }
            ,
            changeParameters() {
                // this.tweetEmbeddings = [];
                // this.loaderVisible = true;
                // // cancel the requests for previous tweets
                // AjaxCaller.cancelRequest();
                // AjaxCaller.resumeRequests();
                // this.updateTopTweets();
            }
            ,
            createEmbeddedUrls: async function (urls) {
                for (let u in urls) {

                    let url = urls[u]._id[0];

                    let endpoint = '/api/embeddings/urls';
                    let data = {
                        url: url
                    };
                    let urlEmbedding = await AjaxCaller.makeInternalCall(endpoint, data);
                    this.loaderVisible = false;

                    this.urls.push(urlEmbedding.data);
                }
            }
            ,
        }
        ,
    })
;

topUrls.initTopUrls();