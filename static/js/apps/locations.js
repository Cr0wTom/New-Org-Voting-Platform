// import {AjaxCaller} from '../modules/ajaxCaller.js';
// import {normalizeValue} from "../modules/utilities.js";
// import {convertRange} from "../modules/utilities.js";


let locations = new Vue({
        el: '#locationsSection',
        data: {},
        methods: {
            initLocations() {

                console.log("Initiating locations...");

                const map = L.map('map').setView([51.505, -0.09], 13);

                L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
                    attribution: '&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
                }).addTo(map);

                L.marker([51.5, -0.09]).addTo(map)
                    .bindPopup('A pretty CSS3 popup.<br> Easily customizable.')
                    .openPopup();
            }
        }
    })
;

// let termsCanvas = document.getElementById('wordcloudCanvas');
// let termsHTML = document.getElementById('wordcloudHTML');
locations.initLocations();

//WordCloud(document.getElementById('wordcloudTerms'), { list: [['foo', 12], ['bar', 6]] } );


