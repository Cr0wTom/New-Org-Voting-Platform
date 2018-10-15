import {AjaxCaller} from '../modules/ajaxCaller.js';
import {convertRange} from "../modules/utilities.js";


let wordclouds = new Vue({
        el: '#wordcloudsSection',
        data: {
            loaderVisible: false,
            canvasVisible: false,
            topic: 'terms',
            language: 'en',
            timestampFrom: '',
            timestampTo: '',
            content: 'terms'
        },
        methods: {
            changeContent() {
                this.setupWordcloudContent();
            },
            changeParameters() {
                this.updateWordclouds();
            },
            initWordclouds() {

                let tenDaysAgoDatetime = moment().subtract(10, 'days');
                let tenDaysAgoTimestamp = moment.unix(tenDaysAgoDatetime);
                let oneDayAgoDatetime = moment().subtract(1, 'days');
                let oneDayAgoTimestamp = moment.unix(oneDayAgoDatetime);


                // Timestamp 1521007715 corresponds to 14 March, the first day we started collecting tweets
                let firstDataDayDatetime = moment.unix(1521007715);
                let todayDatetime = moment();
                const dateRangePicker = $('#dateRange')
                    .daterangepicker({
                        startDate: tenDaysAgoDatetime,
                        endDate: oneDayAgoDatetime,
                        minDate: firstDataDayDatetime,
                        maxDate: todayDatetime,
                        locale: {
                            format: 'DD/MM/YYYY'
                        }

                    })
                    .on('apply.daterangepicker', function () {
                        //wordclouds.updateCalendar(this, fromDate, toDate);
                        let selectedDateFrom = dateRangePicker.data('daterangepicker').startDate;
                        let selectedDateTo = dateRangePicker.data('daterangepicker').endDate;

                        wordclouds.updateCalendar(dateRangePicker, selectedDateFrom, selectedDateTo);

                        wordclouds.updateWordclouds(
                            moment.unix(selectedDateFrom) / 1000,
                            moment.unix(selectedDateTo) / 1000, wordclouds.language, wordclouds.topic
                        )
                    });
                this.updateCalendar(dateRangePicker, tenDaysAgoDatetime, oneDayAgoDatetime);

                this.timestampFrom = tenDaysAgoTimestamp / 1000;
                this.timestampTo = oneDayAgoTimestamp / 1000;
                this.updateWordclouds()
            },
            updateCalendar: function (calendarElement, fromDate, toDate) {
                calendarElement.html('<i class="fas fa-calendar-alt"></i> ' + fromDate.format('MMMM D, YYYY') + ' - ' + toDate.format('MMMM D, YYYY'))
            },
            updateWordclouds: async function () {

                this.loaderVisible = true;
                this.canvasVisible = false;
                let data = {
                    timestamp_from: this.timestampFrom,
                    timestamp_to: this.timestampTo,
                    language: this.language,
                    topic: this.topic
                };

                let endpoint = 'api/wordcloud';
                let response = await AjaxCaller.makeInternalCall(endpoint, data);

                this.loaderVisible = false;
                this.canvasVisible = true;

                this.terms = response.data.terms;
                this.hashtags = response.data.hashtags;
                this.setupWordcloudContent();
            },
            setupWordcloudContent() {
                let list = (this.content === 'terms')
                    ? this.setupWordcloudFormat(this.terms)
                    : this.setupWordcloudFormat(this.hashtags);

                let options = {
                    list: list,
                    gridSize: 5,
                    minFontSize: 10
                };

                WordCloud([termsCanvas, termsHTML], options);
            },
            setupWordcloudFormat(keywords) {

                let formattedWordcloud = [];
                let arr = Object.values(keywords);
                let min = Math.min(...arr);
                let max = Math.max(...arr);

                for (let i in keywords) {
                    if (keywords.hasOwnProperty(i)) {
                        // Adjust each value of the key-value pairs so that the wordcloud looks normal on each initialization.
                        let normalizedRangeValue = convertRange(keywords[i], [min, max], [1, 200]);
                        if (normalizedRangeValue < 10) {
                            normalizedRangeValue += 30;
                        }
                        formattedWordcloud.push([i, normalizedRangeValue])

                    }
                }
                return formattedWordcloud;
            }
        }
    })
;

let termsCanvas = document.getElementById('wordcloudCanvas');
let termsHTML = document.getElementById('wordcloudHTML');
wordclouds.initWordclouds();

//WordCloud(document.getElementById('wordcloudTerms'), { list: [['foo', 12], ['bar', 6]] } );


