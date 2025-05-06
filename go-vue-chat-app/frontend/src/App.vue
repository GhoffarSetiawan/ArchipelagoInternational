<template>
  <div class="app-container">
    <header>
      <h1>Real-Time Chat</h1>
    </header>
    
    <main class="chat-container">
      <!-- User Registration Form -->
      <div v-if="!chatStore.isRegistered" class="registration-container">
        <h2>Enter Your Name</h2>
        <form @submit.prevent="registerUser">
          <input 
            type="text" 
            v-model="username" 
            placeholder="Your Name" 
            required
            class="input-field"
          />
          <button type="submit" class="btn">Join Chat</button>
        </form>
      </div>
      
      <!-- Chat Interface (Two Panels) -->
      <div v-else class="chat-interface">
        <!-- Left Panel - User List -->
        <div class="user-list-panel">
          <h3>Online Users</h3>
          <UserList 
            :users="chatStore.onlineUsers" 
            :selectedUser="chatStore.selectedUser"
            @select-user="selectUser"
          />
        </div>
        
        <!-- Right Panel - Chat Messages -->
        <div class="chat-panel">
          <div v-if="chatStore.selectedUser">
            <h3>Chat with {{ chatStore.selectedUser.name }}</h3>
            <div class="messages-container">
              <MessageList :messages="filteredMessages" />
            </div>
            <MessageInput @send-message="sendMessage" />
          </div>
          <div v-else class="no-selection">
            <p>Select a user to start chatting</p>
          </div>
        </div>
      </div>
    </main>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { useChatStore } from './store/chat'
import UserList from './components/UserList.vue'
import MessageList from './components/MessageList.vue'
import MessageInput from './components/MessageInput.vue'

// Store
const chatStore = useChatStore()

// Local state
const username = ref('')

// Initialize WebSocket connection
onMounted(() => {
  chatStore.initializeWebSocket()
})

// Registration handler
const registerUser = () => {
  if (username.value.trim()) {
    chatStore.registerUser(username.value)
  }
}

// User selection handler
const selectUser = (user) => {
  chatStore.selectUser(user)
}

// Send message handler
const sendMessage = (content) => {
  if (chatStore.selectedUser && content.trim()) {
    chatStore.sendMessage(content)
  }
}

// Filter messages for the selected user conversation
const filteredMessages = computed(() => {
  if (!chatStore.selectedUser) return []
  
  // Get messages between current user and selected user
  return chatStore.messages.filter(msg => 
    (msg.from === chatStore.selectedUser.id && msg.to === chatStore.userId) ||
    (msg.from === chatStore.userId && msg.to === chatStore.selectedUser.id)
  )
})
</script>

<style>
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
  font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
}

body {
  background-color: #f7f9fc;
}

.app-container {
  max-width: 1200px;
  margin: 0 auto;
  padding: 20px;
}

header {
  text-align: center;
  margin-bottom: 30px;
}

header h1 {
  color: #2c3e50;
}

.registration-container {
  max-width: 400px;
  margin: 0 auto;
  padding: 20px;
  background-color: white;
  border-radius: 10px;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
  text-align: center;
}

.registration-container h2 {
  margin-bottom: 20px;
  color: #2c3e50;
}

.input-field {
  width: 100%;
  padding: 12px;
  margin-bottom: 15px;
  border: 1px solid #ddd;
  border-radius: 4px;
  font-size: 16px;
}

.btn {
  background-color: #4c6ef5;
  color: white;
  border: none;
  padding: 12px 20px;
  border-radius: 4px;
  cursor: pointer;
  font-size: 16px;
  transition: background-color 0.3s;
}

.btn:hover {
  background-color: #3b5bdb;
}

.chat-interface {
  display: flex;
  height: 70vh;
  border-radius: 10px;
  overflow: hidden;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
}

.user-list-panel {
  width: 30%;
  background-color: white;
  border-right: 1px solid #eee;
  padding: 15px;
  overflow-y: auto;
}

.user-list-panel h3 {
  margin-bottom: 15px;
  color: #2c3e50;
  padding-bottom: 10px;
  border-bottom: 1px solid #eee;
}

.chat-panel {
  width: 70%;
  background-color: white;
  display: flex;
  flex-direction: column;
}

.chat-panel h3 {
  padding: 15px;
  border-bottom: 1px solid #eee;
  color: #2c3e50;
}

.messages-container {
  flex: 1;
  padding: 15px;
  overflow-y: auto;
  background-color: #f9f9f9;
  height: calc(70vh - 130px);
}

.no-selection {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100%;
  color: #888;
  font-size: 18px;
}
</style>
