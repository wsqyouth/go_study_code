(function (window) {
    "use strict";

    /**
     * Creates a new Model instance and hooks up the server.
     *
     * @constructor
     */
    function Model() { }

    /**
     * Creates a new todo model
     *
     * @param {string} [label] The label of the task
     * @param {function} [callback] The callback to fire after the model is created
     */
    Model.prototype.create = async function (label, callback) {
        label = label || "";
        callback = callback || function () { };

        var newItem = {
            label: label.trim(),
        };

        const resp = await fetch("/todo", {
            method: "POST",
            body: JSON.stringify(newItem),
        });
        const todo = await resp.json();
        callback(todo);
    };

    /**
     * Finds and returns a model in server. If no query is given it'll simply
     * return everything. If you pass in a string or number it'll look that up as
     * the ID of the model to find. Lastly, you can pass it an object to match
     * against.
     *
     * @param {string|number|object} [query] A query to match models against
     * @param {function} [callback] The callback to fire after the model is found
     */
    Model.prototype.read = async function (query, callback) {
        var queryType = typeof query;
        callback = callback || function () { };

        if (queryType === "function") {
            callback = query;
            var todoResponse = await fetch("/todo");
            const todos = await todoResponse.json();
            callback(todos);
        } else if (queryType === "string" || queryType === "number") {
            query = parseInt(query, 10);
            var todoResponse = await fetch(`/todo/${query}`);
            const todo = await todoResponse.json();
            callback(todo);
        } else {
            var todoResponse = await fetch("/todo");
            const todos = await todoResponse.json();
            var filter = todos.filter((t) => t.done == query.completed);
            callback(filter);
        }
    };

    /**
     * Updates a model by giving it an ID, data to update, and a callback to fire when
     * the update is complete.
     *
     * @param {number} id The id of the model to update
     * @param {object} data The properties to update and their new value
     * @param {function} callback The callback to fire when the update is complete.
     */
    Model.prototype.update = async function (id, data, callback) {
        id = parseInt(id, 10);
        const resp = await fetch(`/todo/${id}`, {
            method: "PUT",
            body: JSON.stringify(data),
        });
        const todo = await resp.json();
        callback(todo);
    };

    /**
     * Toggle completed a model by giving it an ID, done to update, and a callback to fire when
     * the update is complete.
     *
     * @param {number} id The id of the model to update
     * @param {function} callback The callback to fire when the update is complete.
     */
    Model.prototype.toggle = async function (id, callback) {
        id = parseInt(id, 10);
        const resp = await fetch(`/todo/${id}/done`, {
            method: "POST",
        });
        const todo = await resp.json();
        callback(todo);
    };

    /**
     * Removes a model from server
     *
     * @param {number} id The ID of the model to remove
     * @param {function} callback The callback to fire when the removal is complete.
     */
    Model.prototype.remove = function (id, callback) {
        id = parseInt(id, 10);
        callback(
            this,
            fetch(`/todo/${id}`, {
                method: "DELETE",
            })
        );
    };

    /**
     * Returns a count of all todos
     */
    Model.prototype.getCount = async function (callback) {
        var todos = {
            active: 0,
            completed: 0,
            total: 0,
        };

        var todoResponse = await fetch("/todo");
        const data = await todoResponse.json();
        data.forEach(function (todo) {
            if (todo.completed) {
                todos.completed++;
            } else {
                todos.active++;
            }

            todos.total++;
        });
        callback(todos);
    };

    // Export to window
    window.app = window.app || {};
    window.app.Model = Model;
})(window);
