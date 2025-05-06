import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import { v4 as uuidv4 } from 'uuid'

export const useChatStore = defineStore('chat', () => {
  // State
  const userId = ref('')
  const username = ref('')
  const isRegistered = ref(false)
  const onlineUsers = ref([])
  const selectedUser = ref(null)
  const messages = ref([])
  const socket = ref(null)
  const isConnected = ref(false)

  // Generate a UUID for this user
  userId.value = uuidv4()

  // Actions
  function initializeWebSocket() {
    // Close any existing connection
    if (socket.value) {
      socket.value.close()
    }

    // Create a new WebSocket connection with the user's ID
    const protocol = window.location.protocol === 'https:' ? 'wss:' : 'ws:'
    const host = window.location.hostname
    const port = '8080' // Backend port
    const wsUrl = `${protocol}//${host}:${port}/ws/${userId.value}`
    
    socket.value = new WebSocket(wsUrl)
    
    // WebSocket event handlers
    socket.value.onopen = () => {
      console.log('WebSocket connection established')
      isConnected.value = true
    }
    
    socket.value.onclose = () => {
      console.log('WebSocket connection closed')
      isConnected.value = false
      
      // Try to reconnect after a delay if we were registered
      if (isRegistered.value) {
        setTimeout(() => {
          initializeWebSocket()
        }, 3000)
      }
    }
    
    socket.value.onerror = (error) => {
      console.error('WebSocket error:', error)
    }
    
    socket.value.onmessage = (event) => {
      try {
        const data = JSON.parse(event.data)
        handleIncomingMessage(data)
      } catch (error) {
        console.error('Error parsing WebSocket message:', error)
      }
    }
  }
  
  function handleIncomingMessage(data) {
    console.log('Received message:', data)
    
    // Handle different message types
    switch (data.type) {
      case 'users':
        // Update the online users list
        try {
          const userList = JSON.parse(data.content)
          // Filter out the current user
          onlineUsers.value = userList.filter(user => user.id !== userId.value)
        } catch (error) {
          console.error('Error parsing user list:', error)
        }
        break
        
      case 'message':
        // Add message to the messages array
        messages.value.push({
          ...data,
          timestamp: new Date().toISOString()
        })
        
        // If this message is from the currently selected user, mark as read
        if (data.from === selectedUser.value?.id) {
          // Implement read status logic if needed
        }
        break
        
      default:
        console.log('Unknown message type:', data.type)
    }
  }
  
  function registerUser(name) {
    if (!isConnected.value) {
      console.error('Cannot register: WebSocket not connected')
      return
    }
    
    username.value = name
    
    // Send registration message
    sendToSocket({
      type: 'register',
      from: userId.value,
      to: '',
      content: name
    })
    
    isRegistered.value = true
  }
  
  function selectUser(user) {
    selectedUser.value = user
  }
  
  function sendMessage(content) {
    if (!selectedUser.value || !isConnected.value) {
      return
    }
    
    const message = {
      type: 'message',
      from: userId.value,
      to: selectedUser.value.id,
      content: content,
      timestamp: new Date().toISOString()
    }
    
    // Send to the server
    sendToSocket(message)
    
    // Add to local messages
    messages.value.push(message)
  }
  
  function sendToSocket(data) {
    if (socket.value && isConnected.value) {
      socket.value.send(JSON.stringify(data))
    } else {
      console.error('Cannot send message: WebSocket not connected')
    }
  }

  // Return store state and methods
  return {
    userId,
    username,
    isRegistered,
    onlineUsers,
    selectedUser,
    messages,
    isConnected,
    initializeWebSocket,
    registerUser,
    selectUser,
    sendMessage
  }
})
