/**
 * Relay Controller
 * Copyright 2015 - Nathan Osman
 */

App = Ember.Application.create();

App.IndexRoute = Ember.Route.extend({
    model: function() {
        return $.get('/api/channels');
    }
})

App.ChannelSwitchComponent = Ember.Component.extend({
    tagName: 'tr',
    click: function() {
        var self = this;
        $.ajax({
            type: 'PUT',
            url: '/api/channels/' + this.get('channel').name,
            contentType: 'application/json',
            data: JSON.stringify({
                state: !this.get('channel').state
            })
        }).then(function(response) {
            self.set('channel', response);
        });
    }
});
