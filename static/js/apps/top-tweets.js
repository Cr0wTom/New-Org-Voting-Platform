import {AjaxCaller} from '../modules/ajaxCaller.js';

let topTweets = new Vue({
        el: '#topTweetsSection',
        delimiters: ['${', '}'],
        data: {
            tweetEmbeddings: [],
            loaderVisible: false,
            loadMoreVisible: false,

            by: 'retweets',
            language: 'en',
            collection: 'terms',
            num: 50
        },
        methods: {
            initTopTweets() {
                this.loaderVisible = true;
                this.loadMoreVisible = false;
                this.updateTopTweets();
            }
            ,
            updateTopTweets: async function () {

                let endpoint = '/api/tweets/top-tweets';
                let data = {
                    by: this.by,
                    language: this.language,
                    collection: this.collection,
                    num: this.num
                };
                let response = await AjaxCaller.makeInternalCall(endpoint, data);
                console.log(response);

                let topTweets = response.data;

                this.tweetEmbeddings = [];
                this.getEmbeddedTweets(topTweets);
            }
            ,
            changeParameters() {
                this.tweetEmbeddings = [];
                this.loaderVisible = true;
                // cancel the requests for previous tweets
                AjaxCaller.cancelRequest();
                AjaxCaller.resumeRequests();
                this.updateTopTweets();
            }
            ,
            getEmbeddedTweets: async function (tweets) {
                for (let t in tweets) {

                    let userId = tweets[t].user.id_str;
                    let tweetId = tweets[t].id_str;

                    let endpoint = '/api/embeddings/tweets';
                    let data = {
                        tweet_id: tweetId,
                        user_id: userId
                    };
                    let tweetEmbedding = await AjaxCaller.makeInternalCall(endpoint, data);
                    this.loaderVisible = false;

                    this.tweetEmbeddings.push(tweetEmbedding.data);

                    const script = document.createElement('script');
                    script.src = 'https://platform.twitter.com/widgets.js';
                    document.head.appendChild(script);
                }
                this.loadMoreVisible = true;
            }
            ,
        }
        ,
    })
;

topTweets.initTopTweets();