<template>
<div id="todo">
  <h2>{{ title }}</h2>
  <p>
  <input type="textbox" v-model.trim="todo"/> <button v-on:click="addTodo">Add</button> <button v-on:click="updateTodo">Upd</button>
  </p>
  <p>
    <ul>
      <li v-for="(todo, index) in todoList" v-bind:key="todo">
        <input type="checkbox"/> {{ todo }} <button v-on:click="deleteTodo(index)">Del</button>
      </li>
    </ul>
  </p>
</div>
</template>

<script>
export default {
  name: 'TodoList',
  data: function () {
    return {
      title: 'Todo!',
      todo: '',
      todoList: []
    }
  },
  methods: {
    addTodo: function () {
      this.todoList.push(this.todo)
    },
    deleteTodo: function (index) {
      this.todoList.splice(index, 1)
    },
    updateTodo: function () {
      alert('upsert DB by using post API')
    }
  }
}
</script>

<style scoped>
ul {
  list-style: none;
}
</style>
