# Archipelago International Technical Test Result

# Basic Questions

## How would you design the logic for this process?
Customer submit an image and it will received by the backend. The backend then handle a workflow for content review by moderator. Moderator can approve, reject , or revert 	  	  content to costumer. If content approved by moderator, customer then can publish it. If content rejected by moderator, customer need to submit new image. If content reverted by 	  moderator, customer need to modify image and its additional information than resubmit it.

## What technologies would you use?
- **PostgreSQL**: Store user data (moderators and customers).
- **Elasticsearch**: Store content metadata and support full-text search.
- **Redis**: Cache frequently requested customer data.
- **AWS S3**: Store uploaded content (images, videos, documents).
- **AWS Cloudfront**: CDN for speeding up content delivery.
- **Go Fiber**: RESTful backend for business logic.
- **Vue 3**: Frontend for web app and admin dashboard.
	
## Do you have any data structure in mind to support this based on your technology of choice to handle those data?
#### Customer
| Field     | Type   | Notes   |
|-----------|--------|---------|
| ID        | UUID   | NOT NULL |
| Name      | string | NOT NULL |
| Email     | string | NOT NULL |
| Password  | string | NULL     |

#### Moderator
| Field     | Type   | Notes   |
|-----------|--------|---------|
| ID        | UUID   | NOT NULL |
| Name      | string | NOT NULL |
| UserCode  | string | NOT NULL |
| UserPass  | string | NOT NULL |

#### ContentIdentity
| Field      | Type   | Notes   |
|------------|--------|---------|
| ID         | UUID   | NOT NULL |
| FileName   | string | NOT NULL |
| CustomerID | UUID   | NOT NULL |
| LastStatus | int    | NOT NULL |

#### ContentStatusHistory
| Field        | Type     | Notes   |
|--------------|----------|---------|
| ID           | UUID     | NOT NULL |
| ContentID    | UUID     | NOT NULL |
| Status       | string   | NOT NULL |
| UserID       | UUID     | NOT NULL |
| OccuredDate  | DateTime | NOT NULL |
| Description  | string   | NOT NULL |

#### ContentMetadata
| Field           | Type     | Notes   |
|------------------|----------|---------|
| ID               | UUID     | NOT NULL |
| ContentID        | UUID     | NOT NULL |
| LastStatus       | int      | NOT NULL |
| Dimension        | string   | NOT NULL |
| FileName         | string   | NOT NULL |
| FileSize         | int      | NOT NULL |
| UploadUserID     | UUID     | NOT NULL |
| UploadTime       | DateTime | NOT NULL |
| UploadRemarks    | string   | NOT NULL |
| ReviewUserID     | UUID     | NULL     |
| ReviewDateTime   | DateTime | NULL     |
| ReviewRemarks    | string   | NULL     |
| PublishDateTime  | DateTime | NULL     |
		

# Database Questions

## Level 1
```sh
SELECT * FROM Customers where Country = 'Germany';
```

## Level 2
```sh
SELECT Count(CustomerID) as Total ,Country FROM Customers GROUP BY Country HAVING Count(CustomerID) >= 5 ORDER BY Count(CustomerID) DESC;
```

## Level 3
```sh
select 
T.CustomerName,
Count(T.OrderID) as Total,
 (
        SELECT TOP 1 X.OrderDate
        FROM Orders X
        WHERE X.CustomerID = T.CustomerID
        ORDER BY X.OrderID ASC
    ) AS FirstOrder,
     (
        SELECT TOP 1 X.OrderDate
        FROM Orders X
        WHERE X.CustomerID = T.CustomerID
        ORDER BY X.OrderID DESC
    ) AS LastOrder
from (
select 
c.CustomerID,
c.CustomerName,
o.OrderID,
o.OrderDate
from 
Customers c 
inner join Orders o on c.CustomerID = o.CustomerID) T
Group By T.CustomerID,T.CustomerName HAVING Count(T.OrderID) >= 5;
```

# Javascript/Typescript Questions
## Level 1 
```sh
function countWordFrequency(text: string): Map<string, number> {
  // Convert text to lowercase
  const lowerText = text.toLowerCase();
  
  // Remove punctuation using regex
  const cleanText = lowerText.replace(/[^\w\s]/g, '');
  
  // Split text into words and filter out empty strings
  const words = cleanText.split(/\s+/).filter(word => word.length > 0);
  
  // Count word frequency
  const frequencyMap = new Map<string, number>();
  
  for (const word of words) {
    const currentCount = frequencyMap.get(word) || 0;
    frequencyMap.set(word, currentCount + 1);
  }
  
  return frequencyMap;
}

// Example usage
function main() {
  const text = "Four One two two three Three three four  four   four";
  const wordFrequency = countWordFrequency(text);
  
  // Sort the result to match expected output
  const sortedWords = Array.from(wordFrequency.entries()).sort((a, b) => a[1] - b[1]);
  
  // Print results
  for (const [word, count] of sortedWords) {
    console.log(`${word} => ${count}`);
  }
}

main();
```

## Level 2

```sh
function delay(ms) {
  return new Promise(resolve => setTimeout(resolve, ms));
}

delay(3000).then(() => alert('runs after 3 seconds'));
```


## Level 3

```sh
// Convert fetchData to return a Promise
function fetchData(url) {
  return new Promise((resolve, reject) => {
    setTimeout(() => {
      if (!url) {
        reject("URL is required");
      } else {
        resolve(`Data from ${url}`);
      }
    }, 1000);
  });
}

// Convert processData to return a Promise
function processData(data) {
  return new Promise((resolve, reject) => {
    setTimeout(() => {
      if (!data) {
        reject("Data is required");
      } else {
        resolve(data.toUpperCase());
      }
    }, 1000);
  });
}

// Using async/await
async function fetchAndProcessData() {
  try {
    // Await the fetch operation
    const data = await fetchData("https://example.com");
    
    // Await the process operation
    const processedData = await processData(data);
    
    // Log the final result
    console.log("Processed Data:", processedData);
    
    return processedData; // In case you need the result elsewhere
  } catch (error) {
    // Handle any errors from either operation
    console.error("Error:", error);
    throw error; // Re-throw if you want calling code to handle it
  }
}

// Execute the async function
fetchAndProcessData();
```

## Level 4
Please visit guide [**here**](https://github.com/GhoffarSetiawan/ArchipelagoInternational/blob/main/go-vue-chat-app/README.md) for running backend and frontend of Chat Application.

# Vue.js

## Explain Vue.js reactivity and common issues when tracking changes.
Vue3 reactivity system is built on top of JavaScript Proxies.A Proxy is an object that wraps another object and allows you to intercept and redefine fundamental operations (e.g., 	  getting or setting a property).Vue3 uses Proxy to wrap your data and reactive() objects, so it can track dependencies and trigger updates automatically. Common issue in Vue2 is 	  Adding new properties is not reactive by default, but in Vue3 it will be fully reactive.

## Describe data flow between components in a Vue.js app
Vue's data flow is top-down way (Parent - Child - Grandchild) and event-based for bottom up communication (child to parrent). The parent component passes data to the child using 	  props.The child sends messages to the parent using this.$emit().

## List the most common cause of memory leaks in Vue.js apps and how they can be solved.
If you start timers or animation loops in components and don’t clear them when the component is destroyed, they keep running and hold onto memory. It can be solved with clearing 	  them in the onBeforeUnmount() lifecycle hook.
Creating DOM elements in JS (e.g., with document.createElement) and appending them manually can leak if not properly cleaned.Avoid manual DOM manipulation. If necessary, clean up 	  with removeChild() or similar in onBeforeUnmount().

## What have you used for state management
Pinia is used for state management. It is Official state management library from Vue team. 

## What’s the difference between pre-rendering and server side rendering?
Pre-rendering and Server-Side Rendering (SSR) are both strategies to deliver HTML to the browser in a more optimized and SEO-friendly way, but they differ in when and how that HTML is generated. Pre-rendering generate HTML at build time, but SSR generate at request time. Pre-rendering server content faster because it basically giving static file to 	  browser, whereas SSR serve content per request. SSR always serve fresh content, whereas pre-rendering serve stale content. Pre-rendering generate static file that it is CDN-ready, whereas SSR does not.


# Website Security Best Practises

1. Implement Best Practice for Authentication and Authorization
	- Use Multi-Factor Authentication (MFA) 
	- Enforce strong password policies (complexity, length, rotation)
	- Implement account lockout after failed attempts
	- Grant minimal access rights needed for job functions
	- Implement just-in-time access for privileged accounts
	- Use separate accounts for administrative and regular tasks
	-
2. Encrypt Sensitive Data
	- Use strong, industry-standard encryption algorithms
	- Proper key management and rotation
	- Implement at-rest and in-transit encryption

3. Secure Network Infrastructure
	- Deploy web application firewalls for public-facing applications
	- Use firewalls and intrusion detection/prevention systems
	- Always use https with valid SSL certificate

4. Secure Development Practices
	- Follow secure coding guidelines (OWASP)
	- Validate all input and sanitize output
	- Use SAST and DAST tools

5. Conduct Regular Security Testing
	- Penetration testing and vulnerability assessments
	- Code security reviews and static analysis
	- Regular security architecture reviews

# Website Performance Best Practises

1. Make sure the backend return response in realtime or near real time period
	- Use non-blocking I/O operations or asynchronous processing
	- Implement message queues in write-heavy process
	- Use Load Balancer to distribute workload

2. Effective Caching Strategy
	- Cache frequently accessed data close to consumers
	- Use CDN to serve static resources
	- Monitor cache hit/miss ratios and optimize accordingly

3. Database Optimization
	- Choose right SQL or NoSQL database for specific use case
	- Consider read replicas for heavy read workloads
	- Implement connection pooling

4. Resource Minification and Bundling
	- Minify JavaScript, CSS, and HTML
	- Bundle related resources
	- Remove unused code and dependencies

5. Good Application Architecture
	- Design for horizontal scalability
	- Implement microservices where appropriate
	- Support auto-scaling in response request surge

# GOLANG

```sh
package main

import (
	"fmt"
	"regexp"
	"strings"
)

func countWords(text string) map[string]int {
	// Convert text to lowercase
	text = strings.ToLower(text)

	// Remove punctuation using regex
	reg := regexp.MustCompile(`[^\w\s]`)
	text = reg.ReplaceAllString(text, "")

	// Split text into words
	words := strings.Fields(text)

	// Count word frequency
	frequency := make(map[string]int)
	for _, word := range words {
		frequency[word]++
	}

	return frequency
}

func main() {
	text := "Four, One two two three Three three four  four   four"

	// Count word frequency
	results := countWords(text)

	// Print results
	fmt.Println("Word Count:")
	for word, count := range results {
		fmt.Printf("%s: %d\n", word, count)
	}
}
```


# .Net
```sh
using System;
using System.Collections.Generic;
using System.Text.RegularExpressions;
using System.Linq;
					
public class Program
{
	public static void Main()
	{
		string text = "Four, One two two three Three three four  four   four";
            
            // Count word frequency
            Dictionary<string, int> results = CountWords(text);
            
            // Print results
            Console.WriteLine("Word Count:");
            foreach (KeyValuePair<string, int> pair in results)
            {
                Console.WriteLine($"{pair.Key}: {pair.Value}");
            }
	}
	
	public static Dictionary<string, int> CountWords(string text)
        {
            // Convert text to lowercase
            text = text.ToLower();
            
            // Remove punctuation using regex
            string pattern = @"[^\w\s]";
            text = Regex.Replace(text, pattern, "");
            
            // Split text into words
            string[] words = text.Split(new char[] { ' ' }, StringSplitOptions.RemoveEmptyEntries);
            
            // Count word frequency
            Dictionary<string, int> frequency = new Dictionary<string, int>();
            foreach (string word in words)
            {
                if (frequency.ContainsKey(word))
                {
                    frequency[word]++;
                }
                else
                {
                    frequency[word] = 1;
                }
            }
            
            return frequency;
        }
}
```


# Tools (Rate yourself 1 to 5)

| Name             | Rate     |
|------------------|----------|
| Git              | 4        |
| Redis            | 4        |
| VCode            | 4        |
| Linux            | 3        |
| AWS EC2          | 3        |
| AWS Lambda       | 2        |
| AWS RDS          | 4        |
| AWS CloudWatch   | 1        |
| AWS S3           | 4        |
| Unit testing     | 4        |
| Kanban Boards    | 3        |
