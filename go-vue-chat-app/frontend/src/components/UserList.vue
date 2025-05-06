<template>
  <div class="user-list">
    <div 
      v-for="user in users" 
      :key="user.id"
      @click="$emit('select-user', user)"
      class="user-item"
      :class="{ 'selected': selectedUser && selectedUser.id === user.id }"
    >
      <div class="user-avatar">
        {{ user.name.charAt(0).toUpperCase() }}
      </div>
      <div class="user-details">
        <div class="user-name">{{ user.name }}</div>
        <div class="user-status">Online</div>
      </div>
    </div>
    <div v-if="users.length === 0" class="no-users">
      No users online
    </div>
  </div>
</template>

<script setup>
// Props and emits
defineProps({
  users: {
    type: Array,
    required: true
  },
  selectedUser: {
    type: Object,
    default: null
  }
})

defineEmits(['select-user'])
</script>

<style scoped>
.user-list {
  display: flex;
  flex-direction: column;
  gap: 10px;
}

.user-item {
  display: flex;
  align-items: center;
  padding: 10px;
  border-radius: 6px;
  cursor: pointer;
  transition: background-color 0.2s;
}

.user-item:hover {
  background-color: #f5f7fb;
}

.user-item.selected {
  background-color: #e9effd;
}

.user-avatar {
  width: 40px;
  height: 40px;
  border-radius: 50%;
  background-color: #4c6ef5;
  color: white;
  display: flex;
  align-items: center;
  justify-content: center;
  font-weight: bold;
  margin-right: 10px;
}

.user-details {
  flex: 1;
}

.user-name {
  font-weight: 500;
  color: #2c3e50;
}

.user-status {
  font-size: 12px;
  color: #10b981;
}

.no-users {
  text-align: center;
  padding: 20px;
  color: #888;
  font-style: italic;
}
</style>
