# Performing transactions
## Introduction
As seeing inside the [FAQs](developer-faq.md#what-is-a-transaction), transactions are the way to alter the current chain state by providing it with the actions to take. Inside this page you will see all the messages that represents the available actions that can be used to edit the chain state.  

## Available messages
Here is the list of currently available [messages](developer-faq.md#what-is-a-message) that you can use while creating transactions for the Desmos chain. 

### Posts
* [`MsgCreatePost`](msgs/staging/posts/create-post.md): allows you to create a new post or a comment for an existing post. 
* [`MsgEditPost`](msgs/staging/posts/edit-post.md): allows you to edit a previously created post message.
* [`MsgAddPostReaction`](msgs/staging/posts/add-post-reaction.md): allows you to add a reaction to an existing post. 
* [`MsgRemovePostReaction`](msgs/staging/posts/remove-post-reaction.md): allows you to remove a reaction from a post.
* [`MsgAnswerPoll`](msgs/staging/posts/answer-poll.md): allows you to answer a post's poll.
* [`MsgRegisterReaction`](msgs/staging/posts/register-reaction.md): allows you to register a reaction.

### Profiles
* [`MsgSaveProfile`](msgs/profiles/save-profile.md): allows you to create or edit an existing profile.
* [`MsgDeleteProfile`](msgs/profiles/delete-profile.md): allows you to delete an existing profile.
* [`MsgRequestDTagTransfer`](msgs/profiles/request-dtag-transfer.md): allows you to ask a transfer for another user `dTag`.
* [`MsgAcceptDTagTransferRequest`](msgs/profiles/accept-dtag-transfer-request.md): allows you to accept a dtag transfer.
* [`MsgRefuseDTagTransferRequest`](msgs/profiles/refuse-dtag-transfer-request.md): allows the `dTag` owner to refuse a transfer request.
* [`MsgCancelDTagTransferRequest`](msgs/profiles/cancel-dtag-transfer-request.md): allows the `dTag` request's sender to cancel his request.

### Relationships
* [`MsgCreateRelationship`](msgs/staging/relationships/create-relationship.md): allows you to create a relationship.
* [`MsgDeleteRelationship`](msgs/staging/relationships/delete-relationship.md): allows you to delete a relationship.
* [`MsgBlockUser`](msgs/staging/relationships/block-user.md): allows you to block a user.
* [`MsgUnblockUser`](msgs/staging/relationships/unblock-user.md): allows you to unblock a user.

### Reports
* [`MsgReportPost`](msgs/staging/reports/report-post.md): allows you to report an existing post.

### Params
* [`EditParamsProposal`](msgs/staging/edit_param_proposal.md): allows you to open a proposal to change profile's params.
