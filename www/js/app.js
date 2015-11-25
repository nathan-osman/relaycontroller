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
        $.ajax({
            type: 'PUT',
            url: '/api/channels/' + this.get('name'),
            contentType: 'application/json',
            data: JSON.stringify({
                state: this.get('state')
            })
        }).then(function(response) {
            this.set('state', response.state);
        });
    }
});
