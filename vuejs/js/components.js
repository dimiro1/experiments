Vue.component('item', {
    props: ['text'],
    template: '#item-template'
});

Vue.component('app', {
    props: ['message'],
    template: '#app-template'
})