Vue.component('item', {
    props: ['text'],
    template: '<li>{{ text }}</li>'
});

Vue.component('app', {
    props: ['message'],
    template: '#app-template',
})