/**
 * Relay Controller
 * Copyright 2015 - Nathan Osman
 */

App = Ember.Application.create();

App.IndexRoute = Ember.Route.extend({
    model: function() {
        return [
            {
                'title': 'Item 1',
                'state': true
            },
            {
                'title': 'Item 2',
                'state': false
            },
            {
                'title': 'Item 3',
                'state': false
            }
        ];
    }
})

App.ChannelSwitchComponent = Ember.Component.extend({
    tagName: 'tr',
    click: function() {
        this.set('state', !this.get('state'));
    }
});
