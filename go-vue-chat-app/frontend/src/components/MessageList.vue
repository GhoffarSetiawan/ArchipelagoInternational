<template>
  <div class="message-list">
    <div 
      v-for="(message, index) in messages" 
      :key="index"
      class="message"
      :class="{ 'message-sent': message.from === chatStore.userId, 'message-received': message.from !== chatStore.userId }"
    >
      <div class="message-content">{{ message.content }}</div>
      <div class="message-time">{{ formatMessageTime(message.timestamp) }}</div>
    </div>
    <div v-if="messages.length === 0" class="no-messages">
      No messages yet. Start the conversation!
    </div>
  </div>
</template>

<script setup>
import { useChatStore } from '../store/chat'

// Store
const chatStore = useChatStore()

// Props
defineProps({
  messages: {
    type: Array,
    required: true
  }
})

// Format timestamp to a readable time
const formatMessageTime = (timestamp) => {
  if (!timestamp) return '';
  
  const date = new Date(timestamp);
  return date.toLocaleTimeString([], { hour: '2-digit', minute: '2-digit' });
}
</script>

<style scoped>
.message-list {
  display: flex;
  flex-direction: column;
  gap: 15px;
}

.message {
  max-width: 70%;
  padding: 10px 15px;
  border-radius: 10px;
  position: relative;
}

.message-sent {
  align-self: flex-end;
  background-color: #4c6ef5;
  color: white;
  border-bottom-right-radius: 2px;
}

.message-received {
  align-self: flex-start;
  background-color: #e9ecef;
  color: #2c3e50;
  border-bottom-left-radius: 2px;
}

.message-content {
  margin-bottom: 4px;
  word-break: break-word;
}

.message-time {
  font-size: 10px;
  opacity: 0.8;
  text-align: right;
}

.no-messages {
  text-align: center;
  padding: 20px;
  color: #888;
  font-style: italic;
}
</style>
