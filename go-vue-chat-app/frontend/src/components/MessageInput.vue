<template>
  <div class="message-input-container">
    <form @submit.prevent="sendMessage">
      <div class="input-wrapper">
        <input 
          type="text" 
          v-model="message" 
          placeholder="Type your message..." 
          class="message-input"
          ref="inputField"
        />
        <button 
          type="submit" 
          class="send-button"
          :disabled="!message.trim()"
        >
          Send
        </button>
      </div>
    </form>
  </div>
</template>

<script setup>
import { ref } from 'vue'

// Emit events to parent
const emit = defineEmits(['send-message'])

// Local state
const message = ref('')
const inputField = ref(null)

// Send message handler
const sendMessage = () => {
  if (message.value.trim()) {
    emit('send-message', message.value)
    message.value = ''
    
    // Focus input for better UX
    if (inputField.value) {
      inputField.value.focus()
    }
  }
}
</script>

<style scoped>
.message-input-container {
  padding: 15px;
  border-top: 1px solid #eee;
}

.input-wrapper {
  display: flex;
  gap: 10px;
}

.message-input {
  flex: 1;
  padding: 12px 15px;
  border: 1px solid #ddd;
  border-radius: 4px;
  font-size: 16px;
  outline: none;
  transition: border-color 0.3s;
}

.message-input:focus {
  border-color: #4c6ef5;
}

.send-button {
  background-color: #4c6ef5;
  color: white;
  border: none;
  padding: 0 20px;
  border-radius: 4px;
  cursor: pointer;
  font-size: 16px;
  transition: background-color 0.3s;
}

.send-button:hover:not(:disabled) {
  background-color: #3b5bdb;
}

.send-button:disabled {
  background-color: #a5b4fc;
  cursor: not-allowed;
}
</style>
