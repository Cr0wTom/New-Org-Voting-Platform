import {AjaxCaller} from '../modules/ajaxCaller.js';

let repository = new Vue({
        el: '#repositorySection',
        delimiters: ['${', '}'],
        data: {
            loaderVisible: false,
            things: []
        },
        methods: {
            initThings() {
                this.loaderVisible = true;
                this.updateThings();
                this.loaderVisible = false;

            },
            updateThings: async function () {

                let endpoint = 'api/repository';
                let data = {
                    source: 'popular'
                };
                let response = await AjaxCaller.makeInternalCall(endpoint, data);

                this.things = response.data.things;
                console.log(this.things);
            }
        }
    })
;

repository.initThings();