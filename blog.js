      async function fetchAndDisplayPosts() {
        try {
          const response = await fetch(
            "https://idleworkshop.com/api/get-posts",
          );

          if (!response.ok) {
            throw new Error(`Error: ${response.statusText}`);
          }

          const data = await response.json();
          console.log(data); // debug

          const posts = data.posts;

          if (!Array.isArray(posts)) {
            throw new Error("Posts is not an array.");
          }

          const postsContainer = document.getElementById("posts-container");
          postsContainer.innerHTML = "";

          for (const post of posts) {
            const postElement = document.createElement("div");
            postElement.classList.add("blog-post");

            const titleElement = document.createElement("h2");
            titleElement.textContent = post.title;

            const dateElement = document.createElement("p");
            dateElement.textContent = `Published: ${new Date(post.date).toLocaleDateString()}`;

            const tagsElement = document.createElement("ul");
            console.log(`Post Title: ${post.title}, Tags:`, post.tags);
            for (const tag of post.tags) {
              const tagElement = document.createElement("li");
              tagElement.textContent = tag;
              tagsElement.appendChild(tagElement);
            };

            const contentElement = document.createElement("md-block");
            contentElement.textContent = post.content;

            postElement.appendChild(titleElement);
            postElement.appendChild(dateElement);
            postElement.appendChild(tagsElement);
            postElement.appendChild(contentElement);

            postsContainer.appendChild(postElement);
          }
        } catch (error) {
          console.error(error);
          const postsContainer = document.getElementById("posts-container");
          postsContainer.innerHTML = "<p>Failed to fetch posts.</p>";
        }
      }
      document.addEventListener("DOMContentLoaded", fetchAndDisplayPosts);

