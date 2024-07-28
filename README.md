##### Basic CI/CD Pipeline using Github Actions on Single Compute Instance

- `CI` Will be triggered on every pr to the main branch. It only runs the test cases.
- `CD` Will be triggered on every push to the main branch. It will deploy the application on the compute instance via ssh.

